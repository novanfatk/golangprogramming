package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeContryToJapanes(address *Address) {
	address.City = "Japanese"
}

func main() {
	address1 := Address{
		City:     "Tegal",
		Province: "Jawa Tengah",
		Country:  "Indonesia",
	}
	address2 := &address1
	address3 := &address1

	address2.City = "Jakarta"
	*address2 = Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	fmt.Println(address1) //tidak berubah jika ada & maka berubah
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "malang"
	fmt.Println(address4)

	alamat := Address{
		City:     "",
		Province: "japan timor",
		Country:  "osaka",
	}
	changeContryToJapanes(&alamat)
	fmt.Println(alamat)
}
