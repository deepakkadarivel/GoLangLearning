package deep

import "fmt"

// PrintVar is exportable because it is capitalized.
func PrintVar() {
	// Myname is exported because it is Capitalized and it belongs to same package
	fmt.Println("My name is : ", Myname)
	// systemname is exported because is belongs to same package
	fmt.Println("System name is : ", systemname)
}
