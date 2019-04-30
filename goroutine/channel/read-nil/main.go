package main

import (
	"fmt"
	"time"
)

func main() {
	var done chan int
	fmt.Println(done == nil)
	go func() {
		time.Sleep(4 * time.Second)
		done <- 1
	}()

	<-done
	fmt.Println("exit")
}
