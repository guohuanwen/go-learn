package singleton

import "sync"

type singleton struct {
	text string
}

func (s *singleton) setData(text string) {
	s.text = text
}

var (
	instance *singleton
	once sync.Once
)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

/*
单例
sync.onces
 */