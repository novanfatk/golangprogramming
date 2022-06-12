package main

import (
	"fmt"
	_ "fmt" //kalo tidak pake gunakan (_)
	"golang-dasar/database"
	_ "golang-dasar/database" //kalo tidak pake gunakan (_)
)

func main() {
	result := database.GetConnet()
	fmt.Println(result)
}
