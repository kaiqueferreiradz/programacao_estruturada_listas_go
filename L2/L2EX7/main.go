package main

import "fmt"

func main() {

	fmt.Println("N M:")
	var n, m, ret, cc int
	fmt.Scanln(&n, &m)

	if n == 0 {
		ret = m
	} else if m == 0 {
		ret = n
	} else {
		if n < m {
			cc = n
		} else {
			cc = m
		}
		//ret = cc
		for i := cc; i >= 0; i-- {

			if n%i == 0 && m%i == 0 {
				ret = i
				i = -1
			}

		}
	}
	fmt.Println(ret)

}
