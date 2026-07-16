package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Numero:")
	var i float64
	fmt.Scanln(&i)

	b := math.Sqrt(i) + (i / 2) + math.Pow(i, i)

	fmt.Println(b)
}
