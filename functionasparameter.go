package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("Hello ", nameFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return " ******* "
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("novan", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
