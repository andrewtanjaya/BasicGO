package main

import "fmt"

func main() {
	// 1
	var nama string
	nama = "Andrew"
	fmt.Println(nama)

	nama = "Andrew Tan"
	//nama = true

	fmt.Println(nama)

	// 2
	var umur = 20
	fmt.Println(umur)

	// 3
	negara := "indonesia"
	fmt.Println(negara)
	negara = "jepang"
	fmt.Println(negara)

	// 4
	var (
		firstname = "Andrew"
		lastname = "Tan"
	)

	fmt.Println(firstname, lastname)

}
