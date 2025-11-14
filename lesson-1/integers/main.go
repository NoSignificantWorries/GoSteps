package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num_str string = "5"

	num, _ := strconv.Atoi(num_str)

	fmt.Println(num * 2)
}
