package testpackage

import "fmt"

// it start with UP letter so it is public
func MyFunction(step int) {

	if step == 1 {
		fmt.Println("step 1")
	} else if step == 2 {
		fmt.Println("step 2")
	} else if step == 3 {
		fmt.Println("step 3")
	} else if step == 4 {
		fmt.Println("step 4")
	} else {
		fmt.Println("step not supported")
	}

}
