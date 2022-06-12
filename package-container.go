package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("novan")
	data.PushBack("java")
	data.PushBack("golang")
	data.PushBack("budi")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
