package main

import "fmt"

func wrapper() func() int {
	x := 32

	return func() int {
		x++
		return x
	}
}

func snack() func() int {
	y := 54

	return func() int {
		y--
		return y
	}
}

func main() {
	increment := wrapper()
	decrement := snack()
	fmt.Println("Value of x is : ", increment())
	fmt.Println("Value of Y is : ", decrement())
}


/*
closure limits the use of variable used by multiple functions
without closure, for functions using that variable
that variable need to be package scope
 */