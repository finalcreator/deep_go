package main

import (
	"fmt"
	"time"
)

func tn(i int) {
	fmt.Println(i)
}

func main() {
	/*	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Hello from goroutine %d\n", i)
		}(i)
	}*/
	go tn(1)
	go tn(2)
	go tn(3)
	go tn(4)
	time.Sleep(time.Duration(6) * time.Second)
}
