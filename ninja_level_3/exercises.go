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

// func main() {
// 	for i := 1; i <= 26; i++ {
// 		fmt.Println(i)
// 		for j := 1; j <= 3; j++ {
// 			fmt.Printf("\t%#U\n", i+64)
// 		}
// 	}
// }

// Exercise 3

// func main() {
// 	year := 1989
// 	for year <= time.Now().Year() {
// 		fmt.Println(year)
// 		year++
// 	}
// }

// Exercise 4

// func main() {
// 	year := 1989
// 	for {
// 		if year <= time.Now().Year() {
// 			fmt.Println(year)
// 			year++
// 		} else {
// 			break
// 		}
// 	}
// }

// Exercise 5

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
