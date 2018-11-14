package main

import (
	"fmt"
)

const (
	d = 1 << iota
	e = iota * 2
	f = iota * 3
)

const (
	a = 1 << iota // a == 1 (iota has been reset)
	b = 1 << iota // b == 2
	c = 1 << iota // c == 4
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
