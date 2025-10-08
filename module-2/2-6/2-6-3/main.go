package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	num_2 := num * num
	num_4 := num_2 * num_2

	fmt.Println(num_4 * num_2)
}
