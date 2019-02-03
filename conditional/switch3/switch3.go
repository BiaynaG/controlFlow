// switch with multiple cases
package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "hello", "bond", "Bond":
		fmt.Println("hello bond or Bond")
	case "bye":
		fmt.Println("this is not the case")
	default:
		fmt.Printf("nothing")
	}
}
