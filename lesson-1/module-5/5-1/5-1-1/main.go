package main

import "fmt"

func main() {
	var size int

	fmt.Scan(&size)

	array := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	for i := 0; i < size; i++ {
		if i%3 == 0 {
			fmt.Print(" ", array[i])
		}
	}
}
