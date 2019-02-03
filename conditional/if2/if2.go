package main

import "fmt"

func main() {
	// scope of x is limited to the if statement
	if x := 10; x == 10 { // use semicolon for two statements on one line
		fmt.Println("001")
	}
	// fmt.Println(x) here will not work, because this is outside the scope of x

}
