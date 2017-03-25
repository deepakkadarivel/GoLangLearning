package main

import "fmt"

func main() {
	isEven := func(n int) (float64, bool) {
		return float64(n) / 2, n%2 == 0
	}
	fmt.Println(isEven(648))
}
