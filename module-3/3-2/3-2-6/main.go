package main

import "fmt"

func main() {
	var age int

	fmt.Scan(&age)

	res := "старость"
	if age <= 13 {
		res = "детство"
	} else if age <= 24 {
		res = "молодость"
	} else if age <= 59 {
		res = "зрелость"
	}

	fmt.Println(res)
}
