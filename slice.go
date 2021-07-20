package main

import "fmt"

func main() {
	// kalau blm tau kapasitasny brapa bs kasi ...
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Juni",
		"July",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice = months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// update array slice jg keupdate
	// slice update array juga keupdate
	months[5] = "July_Updated"
	slice[0] = "Juni_Updated"
	fmt.Println(slice)
	fmt.Println(months)
	var slice2 = months[10:]
	var slice3 = append(slice2, "Bulan_Baru")
	slice3[0] = "Desember_baru"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)
	// make type, len, cap
	slice4 := make([]string, 2, 5)
	slice4[0] = "Andrew"
	slice4[1] = "Wiseson"
	fmt.Println(slice4)
}
