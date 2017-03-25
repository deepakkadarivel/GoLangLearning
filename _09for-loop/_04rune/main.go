package main

import "fmt"

func main() {
	for i := 30; i < 50; i++  {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	fmt.Println("Using formatter")

	for i := 50; i < 100; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

	var text = "HELLO"
	fmt.Println("utf-8 value of text is :", []byte(string(text)))
}
