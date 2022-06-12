package main

import "fmt"

func main() {
	var (
		nilaiAkhir = 80
		absensi    = 75

		lulusUjian   = nilaiAkhir >= 80
		lulusAbsensi = absensi >= 80

		lulus = lulusUjian && lulusAbsensi
	)
	fmt.Println(lulus)
	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)
}
