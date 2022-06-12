package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	utc := time.Date(2000, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parses, _ := time.Parse(layout, "2019-9-10")
	fmt.Println(parses)
}
