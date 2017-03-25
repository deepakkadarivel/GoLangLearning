package main

import "fmt"

func printNum(num []int, callback func(int))  {
		for _, val := range num {
			callback(val)
		}
}

func printMessage(msg string, callback func(string)) {
	callback(msg)
}

func printOddNum(num []int, callback func([]int)) {
	var xs []int
	for _, n := range num {
		if n%2 == 0 {
			continue
		}
		xs = append(xs, n)
		if n > 50 {
			break
		}
	}
	callback(xs)
}

func printEvenNumber(num []int, callback func(int) bool) []int  {
	var xs []int
	for _, n := range num {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	printNum([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	} )

	printMessage("Hello Deep", func(s string) {
		fmt.Println(s)
	})

	var ss []int
	for i := 1; i < 100; i++ {
		ss = append(ss, i)
	}

	printOddNum(ss, func(n []int) {
		fmt.Println(n)
	})

	val := printEvenNumber(ss, func(n int) bool {
		if n%2 == 0 {
			return true
		} else {
			return false
		}
	})
	fmt.Println("Even Numbers : \n", val)
}
