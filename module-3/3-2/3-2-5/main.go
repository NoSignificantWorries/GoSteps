package main

import "fmt"

func summ(n int) int {
	return n/100 + n%100/10 + n%10
}

func main() {
	var num int

	fmt.Scan(&num)

	p1 := num / 1000
	p2 := num % 1000

	if summ(p1) == summ(p2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
