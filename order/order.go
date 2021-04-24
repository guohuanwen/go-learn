package order

import (
	"fmt"
	"sync"
)

//mu.Lock()和mu.Unlock()并不在同一个Goroutine中，所以也就不满足顺序一致性内存模型
func errors()  {
	var mu sync.Mutex
	go func() {
		fmt.Println("hello world")
		mu.Lock()
	}()
	mu.Unlock()
}

//当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，main()函数的阻塞状态驱动后台线程继续向前执行
func fix_error0() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello world")
		mu.Unlock()
	}()
	mu.Lock()
}

/*
对于从无缓存通道进行的接收，发生在对该通道进行的发送完成之前
后台线程<-done接收操作完成之后，main线程的done <- 1发送操作才可能完成（从而退出main、退出程序），而此时打印工作已经完成了
 */
func fix_error1()  {
	done := make(chan int)
	go func() {
		fmt.Println("hello world")
		<-done
	}()
	done <- 1
}

/*
上面的代码虽然可以正确同步，但是对通道的缓存大小太敏感：如果通道有缓存，就无法保证main()函数退出之前后台线程能正常打印了
更好的做法是将通道的发送和接收方向调换一下，这样可以避免同步事件受通道缓存大小的影响
对于带缓存的通道，对通道的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是通道的缓存大小。
虽然通道是带缓存的，但是main线程接收完成是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的
 */
func fix_error2()  {
	done := make(chan int, 1)//带缓存通道
	go func() {
		fmt.Println("hello world")
		done <- 1
	}()
	<-done
}

/*
使用sync.WaitGroup来等待N个线程完成后再进行下一步的同步操作
 */
func doneN() {
	var wg sync.WaitGroup
	for i:=0; i< 10;i++ {
		wg.Add(1)
		go func() {
			fmt.Println("hello world")
			wg.Done()
		}()
	}
	wg.Wait()
}


/*
顺序一致性内存模型
Go语言将其并发编程哲学化为一句口号：“不要通过共享内存来通信，而应通过通信来共享内存。”

 */