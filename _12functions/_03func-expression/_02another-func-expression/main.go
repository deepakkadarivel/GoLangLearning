package main

import "fmt"

func greeting() func() string {
	return func() string {
		return "Hello Deep & Go"
	}
}

func main() {
	greet := greeting()
	fmt.Println(greet())
}
