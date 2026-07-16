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

	if b == 0 {
		return 0
	}
	return a + test(a, b-1)

}
