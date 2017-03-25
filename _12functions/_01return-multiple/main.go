package main

import "fmt"

func getName(fName, lName string) (string, string) {
	return fmt.Sprint(fName, lName), fmt.Sprint(lName, fName)
}

func main() {
	name, reverseName := getName("deepak", "kadarivel")
	fmt.Println("Name : ", name)
	fmt.Println("Reverse Name", reverseName)

}
