package database

import "fmt"

var connect string

func init() {
	connect = "mysql"
	fmt.Println("init dipanggil")
}

func GetConnet() string {
	return connect
}
