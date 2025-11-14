package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	res := ""
	if num < 60 {
		res = "Легкий вес"
	} else if num < 64 {
		res = "Первый полусредний вес"
	} else if num < 69 {
		res = "Полусредний вес"
	}

	fmt.Println(res)
}
