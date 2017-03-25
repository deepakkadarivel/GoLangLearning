package main

import "fmt"

// contact type created using struct
type contact struct {
	name string
	company string
}

func switchByType(x interface{})  {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float32:
		fmt.Printf("float32")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("No type found")
	}
}

func switchForMultipleValues(x int) {
	switch x {
	case 1, 2, 3:
		fmt.Println("this is first block")
	case 4, 5, 6:
		fmt.Println("this is second block")
	default:
		fmt.Println("This doesn't belong to a block.")
	}
}

func main() {

	switchByType(7)
	switchByType("hello")
	var t = contact{name: "deepak", company: "coffeebeans"}
	switchByType(t)
	switchByType(t.name)
	switchByType(34.09)
	switchForMultipleValues(2)
	switchForMultipleValues(4)
	switchForMultipleValues(1)
	switchForMultipleValues(234)

}
