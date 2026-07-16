package main

import (
	"fmt"
)

func main() {
	var s string
	var i, c int
	var ap bool

	fmt.Println("Dados S II CC :")
	fmt.Scanln(&s, &i, &c)

	ap = false
	if s == "M" {
		if i >= 65 && c >= 10 {
			ap = true
		} else if i >= 63 && c >= 15 {
			ap = true
		}
	} else if s == "F" {
		if i >= 63 && c >= 10 {
			ap = true
		} else if i >= 61 && c >= 15 {
			ap = true
		}
	}

	if ap == true {
		fmt.Println("Aposentavel")
	} else {
		fmt.Println("Nao Aposentavel")
	}
}
