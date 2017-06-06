package main

import "fmt"

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	start := 1
	fib := 0
	return func() int {
		fib, start = fib+start, fib
		return fib
	}
}
