package main

import (
	"fmt"
)

func main() {

	var a, b int
	fmt.Println("N:")
	fmt.Scanln(&a, &b)
	fmt.Println(">>>", a)
	fmt.Println(test(a, b))
}

func test(a int, b int) int {

	if b == 1 {
		return a
	}

	if b == a {
		return 1
	}
	fmt.Println(">>>", a)
	return test(a-1, b-1) + test(a-1, b)

}
