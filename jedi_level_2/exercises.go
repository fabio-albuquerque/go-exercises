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

// func main() {
// 	a := 32

// 	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

// 	a = a << 1

// 	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
// }

// Exercise 5

// func main() {
// 	a := `Simon says "hello world!"`

// 	fmt.Println(a)
// }

// Exercise 6

const (
	year1 = 2018 + iota
	year2
	year3
	year4
)

func main() {

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
