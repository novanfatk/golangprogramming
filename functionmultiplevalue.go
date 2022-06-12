package main

import "fmt"

func getFullName() (string, string) {
	return "golang", "java"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
