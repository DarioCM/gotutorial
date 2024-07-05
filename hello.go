package main

import (
	"fmt"
	"gotutorial/testpackage"
)

func main() {

	// var defines variables
	var step int = 1
	testpackage.MyFunction(step)

	var left int = 5
	var right int = 10
	fmt.Println("Sum of", left, "and", right, "is", left+right)

}
