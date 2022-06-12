package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	fmt.Println("argument :", arguments)

	host, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname :", host)
	} else {
		fmt.Println("error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
