package main

import (
	"fmt"
)

// Exercise 1

// func main() {
// 	n := 80
// 	fmt.Printf("%d\t%b\t%#x", n, n, n)
// }

// Exercise 2

// func main() {
// 	b := false
// 	n := 40

// 	if b == false {
// 		b = true
// 	}

// 	if n >= 40 || n <= 40 {
// 		b = true
// 	}

// 	if n < 40 || n > 40 || n != 40 {
// 		b = false
// 	}

// 	fmt.Println(b)
// }

// Exercise 3

// const (
// 	a      = 42
// 	b bool = true
// )

// func main() {

// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// Exercise 4

func main() {
	a := 32

	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	a = a << 1

	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
}
