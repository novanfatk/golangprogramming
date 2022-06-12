package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var contohError error = errors.New("Ups error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
