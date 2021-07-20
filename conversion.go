package main

import "fmt"

func main() {
	angka32 := 500000
	angka64 := int64(angka32)
	angka8 := int8(angka32)

	fmt.Println(angka64)
	fmt.Println(angka32)
	// overflow
	fmt.Println(angka8)

	name := "Andrew"
	a := name[0]
	fmt.Println(string(a))
}
