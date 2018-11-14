package main

import (
	"fmt"
	"sync"
	"time"
)

func addCount(num int, counter chan<- int, wg *sync.WaitGroup) {
	// clear one from the sync group
	defer wg.Done()
	time.Sleep(time.Second * 2)
	counter <- num
}

func main() {
	int_slice := []int{2, 4}
	fmt.Println("length of slice: ", len(int_slice))
	// make the slice buffered using the slice size, so that they can write without blocking
	counter := make(chan int, len(int_slice))

	var wg sync.WaitGroup

	for _, item := range int_slice {
		// add one to the sync group, to mark we should wait for one more
		wg.Add(1)
		go addCount(item, counter, &wg)
	}

	// wait for all goroutines to end
	wg.Wait()

	// close the channel so that we not longer expect writes to it
	close(counter)

	// read remaining values in the channel
	for item := range counter {
		fmt.Println("chan item: ", item)
	}

}
