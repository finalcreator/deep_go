package main

import (
	"bufio"
	"deepgou.com/google_deep_go/functional/fib"
	"fmt"
	"io"
	"strings"
)

type intGen func() uint64

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type FibonacciReader struct {
	Fibo      intGen
	TmpReader *strings.Reader
}

func (fr *FibonacciReader) Read(
	p []byte) (n int, err error) {
	next := fr.Fibo()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	if fr.TmpReader == nil {
		fr.TmpReader = strings.NewReader(s)
	}
	// TODO: incorrect if p is too small!
	return fr.TmpReader.Read(p)
}

func main() {
	//1. origin
	fr := FibonacciReader{fib.Fibonacci(), nil}
	printFileContents(&fr)

}
