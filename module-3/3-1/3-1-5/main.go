package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var p1, p2 string

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	p1 = scanner.Text()
	scanner.Scan()
	p2 = scanner.Text()

	if p1 == p2 {
		fmt.Println("Пароль принят")
	} else {
		fmt.Println("Пароль не принят")
	}
}
