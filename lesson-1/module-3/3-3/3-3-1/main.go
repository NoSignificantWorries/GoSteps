package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	if (-3 <= num && num <= 1) || (5 <= num && num <= 9) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
