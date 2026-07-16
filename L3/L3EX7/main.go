package main

import (
	"fmt"
	"math"
)

func main() {

	var a int
	fmt.Println("N:")
	fmt.Scanln(&a)
	fmt.Println(test(a))
}

func test(a int) int {

	for i := 2.0; i <= 100; i++ {
		for j := 2.0; j <= 100; j++ {
			if math.Pow(i, j) == float64(a) {
				return int(i)
			}
		}
	}
	return 0
}
