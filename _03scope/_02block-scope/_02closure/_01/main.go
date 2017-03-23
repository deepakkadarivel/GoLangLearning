package main

import "fmt"

func main() {
	x := 34
	fmt.Println("value of x : ", x)
	{
		y := "The credit belongs to one who is in the ring."

		fmt.Println("value of y : ", y)
	}
	//not accessible outside closure
	//fmt.Println("value of y outside closure :", y)
}
