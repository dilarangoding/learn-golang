package main

import "fmt"

func main() {
	type NoKTP string

	var ktpJoko NoKTP = "2308230480238402"
	var contoh string = "2323232323232323"
	var contohKTP NoKTP = NoKTP(contoh)
	fmt.Println(ktpJoko)
	fmt.Println(contohKTP)

}
