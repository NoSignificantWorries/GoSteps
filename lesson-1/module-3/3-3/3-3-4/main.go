package main

import "fmt"

func sign(x float64) int {
	if x < 0 {
		return -1
	}
	return 1
}

func main() {
	var x, y float64
	var res int

	fmt.Scan(&x, &y)

	s := sign(x) + sign(y)

	switch s {
	case 2:
		res = 1
	case -2:
		res = 3
	default:
		if x < 0 {
			res = 2
		} else {
			res = 4
		}
	}

	fmt.Println(res)
}
