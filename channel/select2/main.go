package main

import (
	"fmt"
	"time"
)

func main() {
	o := make(chan int)
	c := make(chan int)

	go func() {
		i := 0
		for {

			select {
			case a := <-c: //监听c管道只要一有数据进来 就打印出来
				fmt.Println(a)
			//这里After返回 <-chan Time 也就是监听 <-chan Time这个管道
			//如果超过5秒钟 如果select一直未收到消息 那么 就会给<-chan Time通道发送一个消息
			//每隔5秒就会发送一次
			// case <-time.After(5 * time.Second):
			// 	o <- 0
			// 	fmt.Println(555)
			// 	break //仅仅是跳出select循环并未跳出for循环
			case <-time.After(time.Second):
				i += 1
				fmt.Println(i)
				if i == 5 {
					o <- 0
				}
			}
		}
	}()
	for i := 0; i < 100; i++ {
		c <- i
	}
	<-o //接收消息
}
