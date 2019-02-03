// Create a program that uses a switch statement with the switch expression
// specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "Horseback riding"
	switch favSport {
	case "NoSport":
		fmt.Println("I don't like any sports")
	case "Horseback riding":
		fmt.Println("My favorite sport is horseback riding.")

	}
}
