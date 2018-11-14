package main

import "fmt"
import "time"

func f1(ch chan int) {
	time.Sleep(time.Second * 5)
	ch <- 1
}
func f2(ch chan int) {
	time.Sleep(time.Second * 10)
	ch <- 1
}
func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go f1(ch1)
	go f2(ch2)
	select {
	case <-ch1:
		fmt.Println("The first case is selected.")
	case <-ch2:
		fmt.Println("The second case is selected.")
	}
}
