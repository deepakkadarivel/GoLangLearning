package main

import "fmt"

var x int = 43

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println("Value of x is : ", increment())
	fmt.Println("Value of x is : ", increment())
}


/* closure allows us to limit the scope of variable used by two or more functions
without closure, for two or more functions to have access to same variable
that variable would need to be package scope.
 */