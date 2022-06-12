package main

import "fmt"

func sumALl(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumALl(10, 10, 10, 10, 10)
	fmt.Println(total)

	slice := []int{10, 10, 10, 10, 10}
	total = sumALl(slice...)
	fmt.Println(total)
}
