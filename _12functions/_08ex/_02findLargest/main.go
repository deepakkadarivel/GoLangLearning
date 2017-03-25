package main

import "fmt"

func findLargest(n ...int) int {
	var largest int
	for _, val := range n{
		if val > largest {
			largest = val
		}
	}

	return largest
}

func main() {
	fmt.Println(findLargest(1,2,3,4,5,6,12,24,9,8))
	var xs = []int{23,80,34,12,89,45,67,58}
	fmt.Println(findLargest(xs...))
}
