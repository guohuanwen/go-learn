package subscribe

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

/**
type subscriber chan interface{}
type topicFunc func(v interface{}) bool
可以合并为下面
*/
type (
	subscriber chan interface{}
	topicFunc func(v interface{}) bool
)

type publisher struct {
	m 				sync.RWMutex
	buffer 			int
	timeout			time.Duration
	subscribers 	map[subscriber]topicFunc
}

func newPublisher(publishTimeout time.Duration, buffer int) *publisher {
	return &publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *publisher) subscribeTopic(topic topicFunc) chan interface{}{
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup)  {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
		case sub <- v:
		case <- time.After(p.timeout):
	}
}

func (p * publisher) subscribeAll() chan interface{} {
	return p.subscribeTopic(nil)
}

func (p * publisher) evict(sub chan interface{})  {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

func (p *publisher) publish(v interface{})  {
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *publisher) close()  {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func Test()  {
	fmt.Println("xx")
	p := newPublisher(100 * time.Millisecond, 10)
	defer p.close()
	all := p.subscribeAll()
	golang := p.subscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.publish(" hello world")
	p.publish("hello golang")
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()
	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()
	time.Sleep(3 * time.Second)
}

/*
发布订阅
*/