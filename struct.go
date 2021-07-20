package main

import "fmt"

type Person struct{
	FirstName, LastName string
	Age int
}

func main() {
	var andrew Person
	andrew.FirstName = "Andrew"
	andrew.LastName = "Tanjaya"
	andrew.Age = 20

	fmt.Println(andrew.FirstName)
}
