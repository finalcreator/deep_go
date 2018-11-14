package main

import (
	"fmt"
	"time"
)

func main() {
	mQueue := make(chan string, 2)
	mQueue <- "mike"
	fmt.Println(1)

	mQueue <- "jack"
	fmt.Println(2)

	mQueue <- "lucy"
	fmt.Println(3)

	//close(mQueue)
	for m := range mQueue {
		fmt.Println(m)
	}
	time.Sleep(time.Duration(3) * time.Second)
}
