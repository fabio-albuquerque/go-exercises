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

// func main() {
// 	for i := 10; i <= 100; i++ {
// 		if i%4 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// Exercise 6

// func main() {
// 	b := "james Bond1"
// 	if "james bond" == strings.ToLower(b) {
// 		fmt.Println(true)
// 	} else {
// 		fmt.Println(false)
// 	}
// }

// Exercise 7

// func main() {
// 	b := "james Bond1"
// 	if "james bond" == strings.ToLower(b) {
// 		fmt.Println("true 1")
// 	} else if "james bond1" == strings.ToLower(b) {
// 		fmt.Println("true 2")
// 	} else {
// 		fmt.Println("false")
// 	}
// }

// Exercise 8

// func main() {
// 	switch {
// 	case true:
// 		fmt.Println(true)
// 	case false:
// 		fmt.Println(false)
// 	}
// }

// Exercise 9

// func main() {
// 	favSport := "tEsT"
// 	switch favSport {
// 	case "test":
// 		fmt.Println("test")
// 	case "TEST":
// 		fmt.Println("TEST")
// 	default:
// 		fmt.Printf("Default: %v", favSport)
// 	}
// }

// Exercise 10

func main() {
	fmt.Printf("true && true is: %v\n", true)
	fmt.Printf("true && false is: %v\n", false)
	fmt.Printf("true || true is: %v\n", true)
	fmt.Printf("true || false is: %v\n", true)
	fmt.Printf("!true is: %v\n", true)
}
