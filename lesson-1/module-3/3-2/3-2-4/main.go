package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	n1 := num / 100
	n2 := num % 100 / 10
	n3 := num % 10

	if n1 != n2 && n2 != n3 && n1 != n3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
