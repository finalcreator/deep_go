package main

import (
	"fmt"
	"time"
)

func main() {
	var info chan string
	go func() {
		time.Sleep(1 * time.Second)
		info = make(chan string)
		time.Sleep(1 * time.Second)
		info <- "hello1"
		info <- "hello2"
	}()
	for s := range info {
		fmt.Println(s)
	}
	fmt.Print("Exit")
}
