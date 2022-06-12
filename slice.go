package main

import "fmt"

func main() {

	var month = [...]string{
		"januari",
		"febuari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// month[5] = "diubah"
	// fmt.Println(slice1)

	slice1[0] = "mei lagi"
	fmt.Println(month)

	var slice2 = month[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "novan")
	fmt.Println(slice3)

	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "novan"
	newSlice[1] = "golang"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
