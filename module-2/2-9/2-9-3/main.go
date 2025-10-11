package main

import "fmt"

func main() {
	var F float64

	fmt.Scan(&F)

	fmt.Println(float64(5.0/9.0) * (F - 32))
}
