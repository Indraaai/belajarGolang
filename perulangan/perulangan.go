package main

import "fmt"

func main() {

	// standart perulangan

	for i := 0; i < 5; i++ {
		fmt.Println("nilai =", i)
	}

	// Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	var a = 0

	for a < 8 {
		fmt.Println("angka", a)
		a++
	}

	// Penggunaan Keyword for Tanpa Argumen

	var nilai int = 0

	for {
		fmt.Println("nilai =", nilai)
		nilai++
		if nilai == 5 {
			break
		}
	}

	// Penggunaan Keyword for - range

	// variabel i merupakan index. dan v adalah value dari index itu
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}
	// jika index tidak ingin di munculkan hanya memunculkan value maka var i bisa diganti dengan underscore sehingga tidak ,masalah tidak dgunakan

	// Penggunaan Keyword break & continue
	// continue digunakan untuk memaksa melanjutkan perulangan dan break memakkas perulangan berhenti
	var count = 10
	for i := 1; i < count; i++ {
		if i%2 == 1 {
			continue
		}
		if i >= 8 {
			break
		}
		fmt.Println("nilai ", i)

	}
}
