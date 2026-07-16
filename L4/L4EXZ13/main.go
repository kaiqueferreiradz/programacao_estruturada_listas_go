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

	domi(b)

}

func domi(a [3][3]int) {

	e := [3][3]string{}
	ff := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i-1 >= 0 && a[i-1][j] == 1 {
				ff = ff + 1
			}
			if i+1 < 3 && a[i+1][j] == 1 {
				ff = ff + 1
			}
			if j+1 < 3 && a[i][j+1] == 1 {
				ff = ff + 1
			}
			if j-1 >= 0 && a[i][j-1] == 1 {
				ff = ff + 1
			}
			if i-1 >= 0 && j-1 >= 0 && a[i-1][j-1] == 1 {
				ff = ff + 1
			}
			if i-1 >= 0 && j+1 < 3 && a[i-1][j+1] == 1 {
				ff = ff + 1
			}
			if i+1 < 3 && j+1 < 3 && a[i+1][j+1] == 1 {
				ff = ff + 1
			}
			if i+1 < 3 && j-1 >= 0 && a[i+1][j-1] == 1 {
				ff = ff + 1
			}

			if a[i][j] == 1 {
				e[i][j] = "*"
			} else {
				e[i][j] = fmt.Sprintf("%d", ff)
			}
			ff = 0
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s ", e[i][j])
		}
		fmt.Printf("\n")
	}
}
