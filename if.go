package main

import "fmt"

func main() {
	name := "joko"

	if name == "Novan" {
		fmt.Println("Hello Novan")
	} else if name == "kiki" {
		fmt.Println("Hello kiki")
	} else {
		fmt.Println("hai, kenalan dong")
	}

	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
