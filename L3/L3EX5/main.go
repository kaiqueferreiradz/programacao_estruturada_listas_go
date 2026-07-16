package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Println("N:")
	fmt.Scanln(&a)
	fmt.Println(raiz(a))
}

func raiz(n float64) float64 {

	ret := n / 2
	var ii float64
	ii = 2.0
	for i := 0; i < 20; i++ {
		y := math.Pow(ret, 2) - n
		if y > 0 {
			ret = ret - ret/ii
		} else if y < 0 {
			ret = ret + ret/ii
		} else {
			i = 20
		}
		ii++
	}
	return ret

}
