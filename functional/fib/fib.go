package fib

// 1, 1, 2, 3, 5, 8, 13, ...
func Fibonacci() func() uint64 {
	a, b := 0, 1
	return func() uint64 {
		a, b = b, a+b
		return uint64(a)
	}
}
