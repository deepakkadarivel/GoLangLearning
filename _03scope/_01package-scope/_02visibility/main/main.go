package main

import (
	"fmt"
	"github.com/deepakkadarivel/_03scope/_01package-scope/_02visibility/deep"
)

func main() {
	fmt.Println("Exported value of Myname from Deep is : ", deep.Myname)
	deep.PrintVar()
}
