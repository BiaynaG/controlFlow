package main

import "fmt"

func main() {
	x := 10
	if x == 10 {
		fmt.Println("the value is 10")
	} else if x == 30 {
		fmt.Println("the value is 30")
	} else {
		fmt.Println("the value is not 10")
	}
}
