package main

import (
	"fmt"
)

// Level 1
// func main() {
// 	x := 42
// 	y := "James Bond"
// 	z := true

// 	fmt.Printf("%v\n%v\n%v\n", x, y, z)

// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }

// Level 2
// var x int
// var y string
// var z bool

// func main() {

// 	fmt.Printf("%v\n%v\n%v\n", x, y, z)

// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }

// Level 3
var x = 42
var y = "James Bond"
var z = true

func main() {

	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)

	fmt.Println(s)
}
