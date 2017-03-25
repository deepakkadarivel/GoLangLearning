package main

import "fmt"

func main() {
	var x int = 43

	fmt.Println("Value of x is: ", x)
	fmt.Println("address of x is: ", &x)
	fmt.Printf("decimal address of x is: %d \n", &x)
}
