package consumer

import (
	"fmt"
	"time"
)

func producer(factor int, out chan <- int)  {
	for i := 0 ; ; i++ {
		out <- i * factor
	}
}

func consumer(in <- chan int)  {
	for v := range in {
		fmt.Println(v)
	}
}

func Test() {
	fmt.Printf("xx")
	ch := make(chan int, 64)
	go producer(2, ch)
	go producer(5, ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}

/*
生产消费
*/