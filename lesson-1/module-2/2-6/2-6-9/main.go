package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a, &b, &n)

	a *= n
	a += (b * n) / 100
	b = (b * n) % 100

	fmt.Println(a, b)
}
