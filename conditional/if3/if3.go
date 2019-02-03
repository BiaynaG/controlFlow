// Print all the even numbers up to 100 using for and if loops.

package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
