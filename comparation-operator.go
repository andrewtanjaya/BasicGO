package main

import "fmt"

func main() {
	name1 := "Andrew"
	name2 := "Andrew"

	result := name1 > name2
	fmt.Println(result)

	angka1 := 100
	angka2 := 200

	fmt.Println(angka1 > angka2)
	fmt.Println(angka1 < angka2)
	fmt.Println(angka1 <= angka2)
	fmt.Println(angka1 >= angka2)
	fmt.Println(angka1 != angka2)
	fmt.Println(angka1 == angka2)
}
