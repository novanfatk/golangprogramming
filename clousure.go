package main

import "fmt"

func main() {
	name := "novan"
	counter := 0

	increment := func() {
		name := "budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
