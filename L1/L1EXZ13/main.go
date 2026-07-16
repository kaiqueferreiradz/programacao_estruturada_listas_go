package main

import "fmt"

func main() {
	var d1, m1, a1 int
	var d2, m2, a2 int
	var data bool
	var ig bool

	fmt.Println("Data 1 (DD MM AAAA):")
	fmt.Scanln(&d1, &m1, &a1)

	fmt.Println("Data 2 (DD MM AAAA):")
	fmt.Scanln(&d2, &m2, &a2)

	if a1 > a2 {
		data = true
	} else if a2 > a1 {
		data = false
	} else {

		if m1 > m2 {
			data = true
		} else if m2 > m1 {
			data = false
		} else {
			if d1 > d2 {
				data = true
			} else if d2 > d1 {
				data = false
			} else {
				ig = true
			}
		}

	}

	if ig == true {
		fmt.Println("Iguais")
	} else {
		if data == true {
			fmt.Println("Data 1 Maior")
		} else {
			fmt.Println("Data 2 Maior")
		}
	}
}
