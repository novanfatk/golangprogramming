package main

import "fmt"

func main() {
	name := "novan"

	switch name {
	case "novan":
		fmt.Println("Hello novan")
	case "kiki":
		fmt.Println("Hello kiki")
	case "quenull":
		fmt.Println("Hello quenull")
	default:
		fmt.Println("hai kenalan dong")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}

}
