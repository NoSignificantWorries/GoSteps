package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	num %= 1000

	fmt.Println(num/100 + num%100/10 + num%10)
}
