package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	if num >= 12 {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}
}
