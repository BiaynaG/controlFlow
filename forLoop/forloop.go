package main

import "fmt"

//func main() {
// for initStmt (initializaing a variable); condition (while); postStmt {}
// ++ to add one to the operator
//	for i := 0; i <= 100; i++ {
//		fmt.Println(i)
//	}
//}

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done!")
}

// or the same, using if loop
//func main() {
//	x := 1
//	for {
//		if x > 9 {
//			break
//		}
//		fmt.Println(x)
//		x++
//	}
//	fmt.Println("Done!")
//}
