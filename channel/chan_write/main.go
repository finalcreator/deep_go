package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 6)

	go func() {
		var i int
		for {
			select {
			case ch <- i:
				fmt.Println("ch <- ", i)
			default:
				fmt.Println("channel is full")
				time.Sleep(time.Second * 2)
				fmt.Println("\n- - - - - - - - - - -\n")
			}

			i++
		}
	}()

	for {
		v := <-ch
		fmt.Println(v)
	}
}
