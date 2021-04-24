package _chan

import (
	"fmt"
	"sync"
	"time"
)

/*
先入先出 #

目前的 Channel 收发操作均遵循了先进先出的设计，具体规则如下：
先从 Channel 读取数据的 Goroutine 会先接收到数据；
先向 Channel 发送数据的 Goroutine 会得到先发送数据的权利；
*/

/*
无缓存通道
read write 时机测试
同时准备好，才开始同步数据

4种可能性都有

write prepare
read prepare
write finish
read finish

read prepare
write prepare
read finish
write finish

read prepare
write prepare
write finish
read finish

write prepare
read prepare
read finish
write finish

*/
func Test() {
	var limit = make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//time.Sleep(time.Second)
		fmt.Println("write prepare")
		limit <- 1
		fmt.Println("write finish")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		//time.Sleep(time.Second)
		fmt.Println("read prepare")
		<-limit
		fmt.Println("read finish")
		wg.Done()
	}()
	wg.Wait()
}


/*
缓存通道
 */
func Test1()  {
	var limit = make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("write prepare")
		limit <- 1
		limit <- 1
		limit <- 1
		fmt.Println("write finish")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("read prepare")
		<-limit
		fmt.Println("read finish")
		wg.Done()
	}()
	wg.Wait()
}


//保证协程一定能执行
func Test2() {
	var limit = make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("hello world")
		limit <- 1
	}()
	<-limit
}
