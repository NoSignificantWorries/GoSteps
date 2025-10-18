package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Scan(&n1, &n2, &n3)

	res := 2
	if n1 != n2 && n2 != n3 && n1 != n3 {
		res = 0
	} else if n1 == n2 && n2 == n3 {
		res = 3
	}

	fmt.Println(res)
}
