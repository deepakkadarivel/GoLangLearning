package main

import "fmt"

func foo(n ...int) []int {
	return n
}

func main() {
	fmt.Println(foo(1,2,3))
	var xs = []int{34,43,54,21}
	fmt.Println(foo(xs...))
	fmt.Println(foo())
}
