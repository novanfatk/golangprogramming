package main

import "fmt"

func getGoodBy(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBy
	result := sayGoodBye("Novan")
	fmt.Println(result)
}
