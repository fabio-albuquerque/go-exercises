package main

import (
	"fmt"
)

// const (
// 	a = iota
// 	b
// 	c = "42"
// 	d = iota
// 	e = 19
// 	f
// )

func main() {
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)

	// for i := 33; i <= 122; i++ {
	// 	fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	// }

	test := []int{1, 2, 3, 4, 5}
	for i, v := range test {
		fmt.Println(i, v)
	}
}
