package main

import "fmt"

func main() {
	var a, b, res float64
	var op string

	fmt.Scan(&a, &b, &op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			return
		} else {
			res = a / b
		}
	default:
		fmt.Println("Неверная операция")
		return
	}

	fmt.Println(res)
}
