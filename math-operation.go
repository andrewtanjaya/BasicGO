package main

import "fmt"

func main() {

	result := 10+10
	fmt.Println(result)
	fmt.Println(10*10)
	// augmented assignment
	// a = a + 10
	a := 10
	a += 10
	fmt.Println(a)
	// unary operator
	// a = a+1
	a ++
	fmt.Println(a)
}
