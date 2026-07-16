package main

import (
	"fmt"
)

func main() {

	fmt.Println("N -3x3 -:")

	b := [3][3]int{}
	c := 0
	d := 0

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

	fmt.Println("N:")
	fmt.Scan(&c)
	fmt.Println("N:")
	fmt.Scan(&d)

	iline(b, c)
	icol(b, c)
	idia1(b)
	idia2(b)
	simi(b)
	isub(b, c, d)
	ilico(b, c, d)
	domi(b)

}

func iline(a [3][3]int, c int) {

	cc := 0
	for i := 0; i < 3; i++ {
		cc = cc + a[c][i]
	}
	println(">>>", cc/3)

}

func icol(a [3][3]int, c int) {

	cc := 0
	for i := 0; i < 3; i++ {
		cc = cc + a[i][c]
	}
	println(">>>", cc/3)

}

func idia1(a [3][3]int) {

	cc := 0
	for i := 0; i < 3; i++ {
		cc = cc + a[i][i]
	}
	println(">>>", cc)

}

func idia2(a [3][3]int) {

	cc := 0
	for i := 0; i < 3; i++ {
		cc = cc + a[2-i][i]
	}
	println(">>>", cc)

}

func simi(a [3][3]int) {

	ff := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] != a[j][i] {
				ff = 0
			}
		}
	}
	println(">>>", ff)
}

func isub(a [3][3]int, c int, d int) {

	aux := a[c]
	a[c] = a[d]
	a[d] = aux

	fmt.Println(a)

}

func ilico(a [3][3]int, c int, d int) {

	cc := 0
	for i := 0; i < 3; i++ {
		cc = cc + a[c][i]
	}
	for i := 0; i < 3; i++ {
		if i != c {
			cc = cc + a[i][d]
		}
	}
	println(">>>", cc)

}

func domi(a [3][3]int) {

	ff := 0
	cc := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i-1 >= 0 && a[i][j] > a[i-1][j] {
				ff = ff + 1
			}
			if i+1 < 3 && a[i][j] > a[i+1][j] {
				ff = ff + 1
			}
			if j+1 < 3 && a[i][j] > a[i][j+1] {
				ff = ff + 1
			}
			if j-1 >= 0 && a[i][j] > a[i][j-1] {
				ff = ff + 1
			}
			if i-1 >= 0 && j-1 >= 0 && a[i][j] > a[i-1][j-1] {
				ff = ff + 1
			}
			if i-1 >= 0 && j+1 < 3 && a[i][j] > a[i-1][j+1] {
				ff = ff + 1
			}
			if i+1 < 3 && j+1 < 3 && a[i][j] > a[i+1][j+1] {
				ff = ff + 1
			}
			if i+1 < 3 && j-1 >= 0 && a[i][j] > a[i+1][j-1] {
				ff = ff + 1
			}
			println(a[i][j], "xx..", ff)
			if ((j == 0 && i == 0) || (j == 2 && i == 2) || (j == 2 && i == 0) || (j == 0 && i == 2)) && ff == 3 {
				cc++
			} else if ((j == 0 && i > 0 && i < 2) || (j == 2 && i > 0 && i < 2) || (i == 0 && j > 0 && j < 2) || (i == 2 && j > 0 && j < 2)) && ff == 5 {
				cc++
			} else if ff == 8 {
				cc++
			}
			ff = 0
		}
	}
	println(">>>", cc)
}
