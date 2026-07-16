package main

import (
	"fmt"
)

func main() {

	b := 1.0
	c := 45

	fmt.Println("N -:")
	fmt.Scanln(&b)

	m := [][]int{
		{9, 5, 3, 4, 8, 6, 2, 7, 1},
		{1, 2, 7, 9, 3, 5, 8, 4, 6},
		{6, 8, 4, 7, 1, 2, 9, 3, 5},
		{5, 6, 8, 3, 9, 1, 4, 2, 7},
		{4, 9, 1, 2, 6, 7, 3, 5, 8},
		{3, 7, 2, 8, 5, 4, 1, 6, 9},
		{7, 4, 9, 5, 2, 8, 6, 1, 3},
		{2, 3, 6, 1, 7, 9, 5, 8, 4},
		{8, 1, 5, 6, 4, 3, 7, 9, 2}}

	re := true
	for i := 0; i < 9; i++ {
		cc := 0
		for j := 0; j < 9; j++ {
			cc = cc + m[i][j]
		}
		if cc != c {
			re = false
		}
	}
	for i := 0; i < 9; i++ {
		cc := 0
		for j := 0; j < 9; j++ {
			cc = cc + m[j][i]
		}
		if cc != c {
			re = false
		}
	}

	for i := 0; i < 9; i = i + 3 {

		for j := 0; j < 9; j = j + 3 {

			cc := 0
			for k := 0; k < 3; k++ {
				for h := 0; h < 3; h++ {
					cc = cc + m[k+i][h+j]
					fmt.Printf("%d", m[k+i][h+j])
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n\n")
			//fmt.Println(cc)
			if cc != c {

				re = false
			}

		}
	}

	fmt.Println(re)

}
