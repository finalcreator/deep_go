package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		//close(done)
		done <- 1
	}()
	for i := 0; ; i++ {
		select {
		case <-done:
			return
			// default:
			// 	fmt.Println(i)
		}
		fmt.Println(i)
	}
}
