package main

import (
	"fmt"
)

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func eval2(a, b int, op string) (result int, err error) {
	err = nil
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result, _ = div2(a, b)
	default:
		err = new(error)
	}
	return result, err
}

func main() {
	fmt.Println("Error handling")
	if result, err := eval2(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

}
