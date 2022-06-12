package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//struct function
func (customer Customer) sayHai(name string) {
	fmt.Println("Hello", name, "my nama is", customer)
}

func main() {
	var novan Customer
	novan.Name = "Golang"
	novan.Address = "tegal"
	novan.Age = 20

	novan.sayHai("joko")

	fmt.Println(novan.Name)

	joko := Customer{
		Name:    "java",
		Address: "Java.com",
		Age:     17,
	}

	fmt.Println(joko)

}
