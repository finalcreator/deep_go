package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(time.Second / 2)
	}

	close(ch)

}
