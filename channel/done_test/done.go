package main

import (
	"fmt"
	"net/http/httputil"
	"sync"
)

var counter chan int

func doWork(id int,
	w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		counter <- 1
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(
	id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	wg.Add(20)
	counter = make(chan int, 100)
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
	close(counter)
	var countNum int
	for c := range counter {

		countNum += c
	}
	fmt.Print(countNum)
}

func main() {
	//chanDemo()
	bytes, e := httputil.DumpResponse(nil, true)
	if e != nil {
		println(bytes)
	}
}
