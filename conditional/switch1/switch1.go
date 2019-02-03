package main

import "fmt"

func main() {
	switch {
	case false: //evaluating on the boolean
		fmt.Println("nothing is printed 1")
	case (3 == 4):
		fmt.Println("nothing is printed 2")
	case (3 == 3):
		fmt.Println("printed 1")
		fallthrough // use to keep checking the following cases even if the current one is true
		// do not use this usually
	case (3 == 3):
		fmt.Println("printed 2")
		fallthrough
	case false:
		fmt.Println("nothing is printed 3")
	default:
		fmt.Println("default case")
	}
}
