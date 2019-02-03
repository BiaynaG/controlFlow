// Print the ascii decimal numbers from 33 to 122 using loop
// Print the letter equivalents of those numbers (string)

package main

import "fmt"

func main() {
	x := 33
	for x <= 122 {
		fmt.Printf("%v\t%#x\t%#U\n", x, x, x)
		x++
	}
}
