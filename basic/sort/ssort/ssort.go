package main

import (
	"fmt"
)

var toBeSorted = [...]int{4, 3, 6, 7, 9, 1, 2, 5, 10, 0}

func bubbleSort(input []int) {
	// n is the number of items in our list
	n := len(input)

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if input[min] > input[j] {
				min = j
			}
		}
		fmt.Printf("put min in position %d, by swap: %d <-> %d\n", i, input[i], input[min])
		input[i], input[min] = input[min], input[i]
	}

	fmt.Println(input)
}

func main() {
	fmt.Println("Hello World")
	bubbleSort(toBeSorted[:])

}
