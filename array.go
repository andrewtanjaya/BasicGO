package main

import "fmt"

func main() {
	var names [3]string

	names[0] ="Andrew"
	names[1] = "Wiseson"
	names[2] = "Tanjaya"

	fmt.Println(names[0])

	var ages = [3]int{
		18,
		19,
		20,
	}

	fmt.Println(ages)
	fmt.Println(len(ages))

}
