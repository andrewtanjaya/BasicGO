package main

import "fmt"

func main() {
	orang := map[string]string{
		"nama" : "Andrew",
		"alamat" : "abc",
	}

	orang["telp"] = "01321321"
	fmt.Println(orang["telp"])
	fmt.Println(len(orang))
	delete(orang, "telp")
	fmt.Println(orang["telp"])
}
