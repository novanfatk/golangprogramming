package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 100000
	var nilai64 = int64(nilai32)
	var nilai8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "novan"
	var e = name[0]
	var estring = string(e)

	fmt.Println(name)
	fmt.Println(estring)
}
