// Create a program that uses “else if” and “else”.

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		r := i % 4
		if r != 0 {
			fmt.Printf("When %v is divided by 4, the remainder is %v\n", i, r)
		} else if r == 0 {
			fmt.Printf("When %v is divided by 4, the remainder is 0\n", i)
		} else {
			fmt.Printf("You nailed it")
		}
	}
}
