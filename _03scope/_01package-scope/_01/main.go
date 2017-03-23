package main

import "fmt"

var x int = 1

func main()  {
	fmt.Println("Value of X : ", x)
	foo()
}

func foo() {
	fmt.Println("value of x in foo X :", x)
}
