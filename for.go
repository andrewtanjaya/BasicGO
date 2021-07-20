package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10{
		fmt.Println(counter)
		counter++
	}

	// init statement, post statement
	for count:=1 ; count<=10 ; count++{
		fmt.Println(count)
	}

	// iterasi map, array, slice
	names := []string{"andrew", "andi", "budi"}
	for index, name := range names{
		fmt.Println("index ",index, " = ", name)
	}
}
