package main

import "fmt"

func main() {
	x := 45
	fmt.Println("Value of x from declaration block : ", x)
	foo()
}

func foo() {
	// This cannot be accessed because of block scopoe
	//fmt.Println("value of x from another block: ", x)
}
