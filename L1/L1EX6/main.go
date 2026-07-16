package main

import (
	"fmt"
	"math"
)

func main() {

	var a float64
	var b float64
	var c float64

	fmt.Println("3 Lados:")
	fmt.Scanln(&a, &b, &c)

	s := (a + b + c) / 2

	r := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	fmt.Println(r)
}
