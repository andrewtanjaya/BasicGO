package main

import "fmt"

func main() {
	// 1
	const nama = "Andrew"
	const jenis_kelamin = "pria"

	// error
	//nama = "andi"
	//jenis_kelamin = "perempuan"

	fmt.Println(nama)
	fmt.Println(jenis_kelamin)

	// 2
	const (
		negara = "Indonesia"
		tinggi = 180.5
	)

	fmt.Println(negara)
	fmt.Println(tinggi)
}
