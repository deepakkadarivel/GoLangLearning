package main

import "fmt"

func main() {
	var a = 0

	fmt.Println("value of a is : ", a)
	fmt.Println("address of a is : ", &a)

	// reference
	var x *int
	x = &a

	fmt.Println("address of x is : ", x)
	fmt.Println("value of x is : ", *x)

	// dereference
	*x = 23

	fmt.Println("address of x is : ", x)
	fmt.Println("value of x is : ", *x)
	fmt.Println("address of a is : ", &a)
	fmt.Println("value of a is : ", a)
}
