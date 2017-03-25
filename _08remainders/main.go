package main

import "fmt"

func main() {
	x := 13 % 3

	if x == 1 {
		fmt.Println("ODD")
	} else {
		fmt.Println("EVEN")
	}

	for i := 1; i < 10; i++ {
		if i % 2 == 1 {
			fmt.Println("ODD")
		} else {
			fmt.Println("EVEN")
		}
	}
}
