package main

import "fmt"

func main() {
	type nikah bool

	var andrewNikah nikah
	andrewNikah = true

	fmt.Println(andrewNikah)
}
