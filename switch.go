package main

import "fmt"

func main() {
	lulus := false

	switch lulus {
		case false : fmt.Println("Smangat :)")
		case true: fmt.Println("Selamat anda lulus")
	}
	// short statement
	switch lulused := false ; lulused {
	case false : fmt.Println("Smangat :)")
	case true: fmt.Println("Selamat anda lulus")
	}
	// tanpa expression
	switch  {
	case lulus ==false : fmt.Println("Smangat :)")
	case lulus == true: fmt.Println("Selamat anda lulus")

	}
}
