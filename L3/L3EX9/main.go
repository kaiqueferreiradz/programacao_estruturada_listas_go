package main

import (
	"fmt"
)

func main() {

	var a, b int
	fmt.Println("N:")
	fmt.Scanln(&a, &b)
	fmt.Println(test(a, b))
}

func test(a int, b int) int {

	ret := 0
	for i := 1; i <= a; i++ {
		ret = ret + b
	}
	return ret
}
