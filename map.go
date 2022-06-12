package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "novan",
		"address": "Tegal",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "codino"
	book["ups"] = "salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
