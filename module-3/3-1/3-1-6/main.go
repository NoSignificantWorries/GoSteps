package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	res := 0
	if num < 0 {
		res = -1
	} else if num > 0 {
		res = 1
	}

	fmt.Println(res)
}
