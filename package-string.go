package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("novan java", "novan"))
	fmt.Println(strings.Split("novan java golang", " "))
	fmt.Println(strings.ToUpper("novan"))
	fmt.Println(strings.ToLower("NOVAN"))
	fmt.Println(strings.Trim("     novan    ", " "))
	fmt.Println(strings.ReplaceAll("novan novan novan novan", "novan", "bido"))
}
