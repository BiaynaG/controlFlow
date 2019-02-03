package main

import "fmt"

func main() {
	switch "Bond" { // evaluating on the value
	case "Nice":
		fmt.Println("this is not nice")
	case "Bond":
		fmt.Println("This is Bond, James Bond!")
	default:
		fmt.Println("default case")
	}
}
