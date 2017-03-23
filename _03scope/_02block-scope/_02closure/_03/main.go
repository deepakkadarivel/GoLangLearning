package main

import "fmt"

func main() {
	var x int = 43

	increment := func() int {
		x ++
		return x
	}

	fmt.Println("Value of x is : ", increment())
	fmt.Println("value of x is : ", increment())

}

/* closure helps to limit the scope of variable used by multiple functions
without closure, for two or more functions having access to same variable
that variable would need to be package scope

anonymous function
a function without a name

function expression
assigning a function to a variable
 */
