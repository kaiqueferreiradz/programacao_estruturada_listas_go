package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Println("N:")
	fmt.Scanln(&a)
	fmt.Println(">>>", a)
	fmt.Println(test(a))
}

func test(a int) int {

	if a <= 0 {
		return 0
	}
	fmt.Println(">>>", a)
	return a + test(a-1)

}
