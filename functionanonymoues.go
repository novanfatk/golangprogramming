package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked")
	} else {
		fmt.Println("welcome", name)
	}
}

// func blacklisAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklisRoot(name string) bool {
// 	return name == "root"
// }
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("novan", blacklist)

	registerUser("novan", func(name string) bool {
		return name == "root"
	})
}
