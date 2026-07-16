package main

import (
	"fmt"
	"math"
)

func main() {

	b := 1.0
	t := ""

	fmt.Println("N -:")
	fmt.Scanln(&b)
	fmt.Println("TXT -:")
	fmt.Scanln(&t)

	tt := []rune(t)

	c := int(math.Sqrt(b))

	m := [][]string{}

	k := 0

	for i := 0; i < c; i++ {
		aux := []string{}
		for j := 0; j < c; j++ {
			aux = append(aux, string(tt[k]))
			k++
		}
		m = append(m, aux)
	}

	for i := 0; i < c; i++ {

		for j := 0; j < c; j++ {
			if m[j][i] != "*" {
				fmt.Printf("%s", m[j][i])
			}
		}
		fmt.Printf(" ")

	}

}
