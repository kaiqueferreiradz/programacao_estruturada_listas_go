package main

import "fmt"

func main() {

	fmt.Println("N:")
	var n int
	fmt.Scanln(&n)
	b, c, d, e := 0, 0, 0, 0

	for i := 1; i <= n; i++ {
		fmt.Println("N", i, ":")
		var a int
		fmt.Scanln(&a)
		if a <= 25 {
			b++
		} else if a <= 50 {
			c++
		} else if a <= 75 {
			d++
		} else if a <= 100 {
			e++
		}
	}

	fmt.Println("Intervalo [0..25]:", b)
	fmt.Println("Intervalo [26..50]:", c)
	fmt.Println("Intervalo [51..75]:", d)
	fmt.Println("Intervalo [76..100]:", e)

}
