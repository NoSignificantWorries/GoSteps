package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	fmt.Printf("Следующее за числом %d число: %d\nДля числа %d предыдущее число: %d\n", num, num+1, num, num-1)
}
