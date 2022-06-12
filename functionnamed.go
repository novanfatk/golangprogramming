package main

import "fmt"

func getFullNames() (firstName string, middleName string, lastName string) {
	firstName = "novan"
	middleName = "java"
	lastName = "golang"
	return
}

func main() {
	a, b, c := getFullNames()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
