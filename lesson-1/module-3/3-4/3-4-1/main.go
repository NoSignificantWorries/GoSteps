package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	var res int = 31
	switch num {
	case 2:
		res = 29
	case 4:
		res = 30
	case 6:
		res = 30
	case 9:
		res = 30
	case 11:
		res = 30
	default:
		res = 31
	}

	fmt.Println(res)
}
