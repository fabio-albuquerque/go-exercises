package main

import (
	"fmt"
)

// Exercise 1

// func main() {
// 	for i := 1; i <= 10000; i++ {
// 		fmt.Println(i)
// 	}
// }

// Exercise 2

func main() {
	for i := 1; i <= 26; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i+64)
		}
	}
}
