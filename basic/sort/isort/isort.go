package main

import (
	"fmt"
)

var toBeSorted = [...]int{8, 4, 3, 6, 7, 9, 1, 2, 5, 10, 0}

func bubbleSort(input []int) {
	// n is the number of items in our list
	n := len(input)

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if input[j] > input[j-1] {
				break
			}
			fmt.Printf("swap: input[%d]:%d <-> input[%d]:%d\n", j, input[j], j-1, input[j-1])
			input[j], input[j-1] = input[j-1], input[j]
		}
	}

	fmt.Println(input)
}

func main() {
	fmt.Println("Hello World")
	bubbleSort(toBeSorted[:])

}
