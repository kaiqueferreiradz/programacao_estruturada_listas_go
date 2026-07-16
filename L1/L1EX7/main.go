package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Raio:")

	var r float64

	fmt.Scanln(&r)

	a := math.Pi * math.Pow(r, 2)

	p := 2 * math.Pi * r

	fmt.Printf("A = %f , P = %f", a, p)

}
