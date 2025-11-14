package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	var res string = ""
	switch {
	case num == 12 || num == 1 || num == 2:
		res = "Зима"
	case num >= 3 && num <= 5:
		res = "Весна"
	case num >= 6 && num <= 8:
		res = "Лето"
	default:
		res = "Осень"
	}

	fmt.Println(res)
}
