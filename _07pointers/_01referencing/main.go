package main

import "fmt"

func main() {
	var b int = 24

	fmt.Println("value of b is: ", b)
	fmt.Println("address of b is: ", &b)

	var x *int

	x = &b

	fmt.Println("address of x is: ", x)


}
