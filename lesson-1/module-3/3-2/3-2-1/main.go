package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	if num < 1000 && num > 99 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
