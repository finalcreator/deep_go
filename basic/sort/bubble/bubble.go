package main

import (
	"fmt"
)

var toBeSorted = [...]int{6, 3, 7, 9, 1, 2, 5, 10, 0}

func bubbleSort(input []int) {
	// n is the number of items in our list
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				fmt.Printf("Swapping: %d <-> %d\n", input[i], input[i+1])
				// swap values using Go's tuple assignment
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	fmt.Println(input)
}

func main() {
	fmt.Println("Hello World")
	bubbleSort(toBeSorted[:])

}
