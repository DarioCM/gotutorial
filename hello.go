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

	var name string = "Dario"
	fmt.Println("Hello", name, "!")

	var isHappy bool = true
	fmt.Println("Are you happy?", isHappy)

}
