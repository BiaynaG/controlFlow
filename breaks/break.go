// Printing only even numbers until 100
// Using break and continue
package main

import "fmt"

func main() {
	x := 0
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 { // modulus, gives the remainder
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("Done")
}
