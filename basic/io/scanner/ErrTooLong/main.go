package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	buf := make([]byte, 10)
	input := strings.Repeat("x", 20)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Buffer(buf, 21)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}
