package main

import "fmt"

func main() {

	if umur := 50;umur <= 1{
		fmt.Println("Masi Bayi")
	}else if umur <= 5 {
		fmt.Println("Balita")
	}else if umur <= 17 {
		fmt.Println("Remaja")
	}else if umur <= 40 {
		fmt.Println("Dewasa")
	}else {
		fmt.Println("Udah tua")
	}
	//fmt.Println(umur)
}
