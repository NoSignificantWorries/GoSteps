package main

import "fmt"

func main() {
	var minutes int

	fmt.Scan(&minutes)

	hours := minutes / 60
	clear_minutes := minutes % 60

	fmt.Printf("%d мин - это %d час %d минут.\n", minutes, hours, clear_minutes)
}
