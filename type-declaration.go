package main

import "fmt"

func main() {
	type NoKtp string

	type Merried bool

	var noKtpKu NoKtp = "1246346372"
	fmt.Println(noKtpKu)

	var merriedStatus Merried = true
	fmt.Println(merriedStatus)
}
