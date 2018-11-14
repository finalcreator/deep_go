package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "foo   bar   你好吗ok 非常好   baz\t 1\t2  3"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
