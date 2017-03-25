package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for j :=0; j < 4; j++  {
			fmt.Printf("i - %d : j - %d\n", i, j)
		}
	}
}
