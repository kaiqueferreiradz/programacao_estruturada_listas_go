package main

import "fmt"

func main() {

	var n int
	fmt.Println("N:")
	fmt.Scanln(&n)

	k := n
	w := true
	for w {
		p := true
		for i := 2; i < k; i++ {
			if k%i == 0 {
				p = false
			}
		}
		if p {
			fmt.Println(k)
			w = false
		}
		k++
	}

	for j := n; j > 0; j-- {
		p := true
		for i := 2; i < j; i++ {
			if j%i == 0 {
				p = false
			}
		}
		if p {
			fmt.Println(j)
			j = 0
		}
	}

}
