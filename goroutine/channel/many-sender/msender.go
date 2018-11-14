package main

import (
	"fmt"
	_ "sync"
	"time"
)

var mchan chan int = make(chan int, 100)

func feed(id int, mchan chan int, duration float32) {
	i := id
	for {
		time.Sleep(time.Duration(duration) * time.Second)
		fmt.Printf("-> mchan%d got i:%d\n ", id, i)
		mchan <- i
		i++

	}
}

func main() {
	go feed(1000, mchan, 2)
	go feed(2000, mchan, 4)
	go func() {
		for num := range mchan {
			// time.Sleep(time.Duration(1) * time.Second)
			fmt.Println(" *** New message: ", num)

		}
	}()
	ticker := time.NewTicker(time.Second * 1)
	for elem := range ticker.C {
		fmt.Println(" ---- ", elem, " ---- ")
	}
}
