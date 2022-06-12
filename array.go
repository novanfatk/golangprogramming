package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "novan"
	names[1] = "golang"
	names[2] = "java"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	names[0] = "nika"
	fmt.Println(names[0])
	fmt.Println(len(names))

	var values = [3]int{
		90, 80, 70,
	}
	fmt.Println(values)
	fmt.Println(values[0])
}
