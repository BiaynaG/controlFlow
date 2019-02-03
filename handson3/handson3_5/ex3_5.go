// Print out the sentence "When X is divided by 4, the remainder is X\n",
// for all the non-zero remainders (modulus) which are found for each number between
// 10 and 100 when it is divided by 4.

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		r := i % 4
		if r != 0 {
			fmt.Printf("When %v is divided by 4, the remainder is %v\n", i, r)
		}
	}
}
