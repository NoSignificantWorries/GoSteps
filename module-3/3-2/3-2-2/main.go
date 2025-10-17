package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	if num < 17 && num > -1 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}
