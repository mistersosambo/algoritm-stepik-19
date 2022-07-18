package main

import "fmt"

func main() {
	var a, z int
	var m []int
	fmt.Scan(&a)
	fmt.Scan(&z)
	for ; a != 0; a /= 10 {
		r := a % 10
		if r != z {
			m = append(m, r)
		}
	}
	for i := len(m) - 1; i >= 0; i-- {
		fmt.Print(m[i])
	}
}
