package limitMaxThread

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	//获取goroot目录：
	fmt.Println("GOROOT-->",runtime.GOROOT())
	//获取操作系统
	fmt.Println("os/platform-->",runtime.GOOS) // GOOS--> darwin，mac系统
	//获取逻辑cpu的数量
	fmt.Println("逻辑CPU的核数：",runtime.NumCPU())
	//设置go程序执行的最大的：[1,256]
	n := runtime.GOMAXPROCS(runtime.NumCPU() / 4)
	fmt.Println(n)
}


func Test()  {
	work := make(map[string]func())
	work["1"] = func() {
		fmt.Println("hello 1")
	}
	work["2"] = func() {
		fmt.Println("hello 2")
	}
	work["3"] = func() {
		fmt.Println("hello 3")
	}
	work["4"] = func() {
		fmt.Println("hello 4")
	}
	work["5"] = func() {
		fmt.Println("hello 5")
	}
	work["6"] = func() {
		fmt.Println("hello 6")
	}

	var wg sync.WaitGroup
	for _, w := range work {
		wg.Add(1)
		go func() {
			w()
			wg.Done()
		}()
		wg.Wait()
	}

}
