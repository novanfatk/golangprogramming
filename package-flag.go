package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "put your host")
	user := flag.String("user", "root", "put your user")
	pswd := flag.String("password", "root", "put your user")

	flag.Parse()
	fmt.Println(*host, *user, *pswd)
}
