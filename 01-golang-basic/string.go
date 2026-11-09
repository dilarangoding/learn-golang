package main

import "fmt"

func main() {
	firstName := "Joko"
	lastName := "Widodo"

	// String dapat digabungkan menggunakan operator +.
	fullName := firstName + " " + lastName
	fmt.Println("Hello " + fullName)

	// len menghitung jumlah byte, bukan jumlah karakter.
	fmt.Println("Jumlah byte:", len(fullName))

	// Mengakses string dengan indeks menghasilkan byte.
	fmt.Println("Byte pertama:", fullName[0])

	// String UTF-8 dapat menggunakan lebih dari satu byte untuk satu karakter.
	text := "é"
	fmt.Println("Byte pada é:", len(text))
	for _, character := range text {
		fmt.Printf("Karakter: %c\n", character)
	}
}
