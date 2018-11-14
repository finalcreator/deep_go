package main

import (
	"fmt"
)

type Settable interface {
	SetVal(val int)
}

func (c *check) SetVal(val int) {
	fmt.Println("in setvale")
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", *c)
	c.Val = val
}

func some(te Settable) {
	te.SetVal(20)
}

type check struct {
	Val int
}

func main() {
	a := check{Val: 100}
	fmt.Println(a)
	fmt.Println(&a)
	//p := &a
	//some(&a)
	a.SetVal(99)
	fmt.Println(&a)
	fmt.Println(a)
}
