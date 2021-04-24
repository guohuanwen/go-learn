package main

import (
	_chan "learn/chan"
	"log"
)

func init()  {
	
}

func main()  {
	defer log.Println("exit")
	//生产消费模型
	//consumer.Test();
	//发布订阅模型
	//subscribe.Test()
	//limitMaxThread.Test()
	_chan.Test()
}