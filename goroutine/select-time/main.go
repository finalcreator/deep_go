package main

import (
	"fmt"
	"time"
)

func main() {

	sig := func(after time.Duration) <-chan interface{} { // <1>
		c := make(chan interface{})
		go func() {
			defer close(c)
			defer fmt.Printf("%v past", after)
			time.Sleep(after)
		}()
		return c
	}

	select {
	case <-sig(1 * time.Second):
	case <-sig(5 * time.Second):
	case <-sig(2 * time.Second):
	case <-sig(2 * time.Minute):
	case <-sig(2 * time.Hour):
	}
}
