package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Merried() {
	man.Name = "mr. " + man.Name
	//fmt.Println("di method", man.Name)
}

func main() {
	dev := Man{"novan"}
	dev.Merried()

	fmt.Println(dev.Name)
}
