package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib_pair := []int{0, 1}
	return func() int {
		fib_pair[0], fib_pair[1] = fib_pair[1], fib_pair[0] + fib_pair[1]
		return fib_pair[0]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

