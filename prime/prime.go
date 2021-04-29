package prime

import "fmt"

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i:=2 ; ; i++ {
			fmt.Printf("generate %v \n", i)
			ch <- i
		}
	}()
	return ch
}

func primeFilter(in <- chan int ,prime int) chan int {
	fmt.Printf("primeFilter prime %v \n", prime)
	out := make(chan int)
	go func() {
		for {
			i := <-in;
			fmt.Printf("primeFilter in %v  prime %v \n", i, prime)
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func Test() {
	ch := generate()
	for i:= 0; i< 10;i++ {
		fmt.Printf("第%v次循环: \n", i + 1)
		prime := <- ch
		fmt.Printf("第%v个: %v \n", i + 1, prime)
		ch = primeFilter(ch, prime)
	}
}

/*
暴力素数筛
 */
