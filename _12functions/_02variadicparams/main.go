package main

import "fmt"

func average(x ...float64) float64  {
	fmt.Println("value of x is : ", x)
	fmt.Printf("Type of x is : %T \n", x)

	var total float64
	for _, val := range x {
		total += val
	}

	return total / float64(len(x))
}

func main() {
	//n := average(12,34,45,65,76,89)

	data := []float64{123,345,54,32,12}
	n := average(data...)
	fmt.Println("Avarage value is : ", n)
}
