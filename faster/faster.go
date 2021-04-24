package faster

import (
	"fmt"
	"time"
)

func searchByBing(text string) string  {
	time.Sleep(2 * time.Second)
	return "bing golang"
}

func searchByGoogle(text string) string  {
	time.Sleep(1 * time.Second)
	return "google golang"
}


func Test()  {
	ch := make(chan string, 32)
	go func() {
		ch <- searchByBing("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	fmt.Println(<-ch)
}
/*
快先返回
*/
