package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("N:")
	var n float64
	fmt.Scanln(&n)

	var f float64
	m := n / 2
	g := 2.0
	ex := false
	for i := 0; i < 20; i++ {
		f = math.Pow(m, 2) - n

		fmt.Println(i, "__", f, "-", m)

		if f > 0 {
			m = m - m/g
		} else if f < 0 {
			m = m + m/g
			ex = true
		} else {
			i = 20
		}

		if ex {
			g = g + 10.0
		}

	}

}
