package main

import (
	"fmt"
)

func main() {

	fmt.Println("N -3x3 -:")

	b := [3][3]int{}

	fmt.Scan(&b[0][0])
	fmt.Scan(&b[0][1])
	fmt.Scan(&b[0][2])
	fmt.Scan(&b[1][0])
	fmt.Scan(&b[1][1])
	fmt.Scan(&b[1][2])
	fmt.Scan(&b[2][0])
	fmt.Scan(&b[2][1])
	fmt.Scan(&b[2][2])

	fmt.Println(b)

	ret := 1
	n := 0
	m := 0
	for i := 0; i < 3; i++ {
		n = n + b[0][i]
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m = m + b[i][j]
		}
		if m != n {
			ret = 0
		}
		m = 0
	}

	m = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m = m + b[j][i]
		}
		if m != n {
			ret = 0
		}
		m = 0
	}

	m = 0
	for i := 0; i < 3; i++ {
		m = m + b[i][i]
	}
	if m != n {
		ret = 0
	}
	m = 0

	m = 0
	for i := 0; i < 3; i++ {
		m = m + b[i][2-i]
	}
	if m != n {
		ret = 0
	}
	m = 0

	fmt.Println(ret)

}
