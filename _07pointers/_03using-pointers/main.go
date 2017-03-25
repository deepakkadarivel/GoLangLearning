package main

import "fmt"

func foo(x *int) {
	*x = 0
}

func main() {
	var x = 2

	fmt.Println("value of x is : ", x)

	foo(&x)

	fmt.Println("value of x after foo is : ", x)
}
