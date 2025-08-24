package main

import "fmt"

// format penulisan array adalah = varaiabel = , julah panjang  array [], tipe data arraynya, baru isi datanya dalam kurng kurawal {}

func main() {

	var fruits = [4]string{"apel", "mangga", "jambu", "pepaya"}
	fmt.Println(fruits)

	//  Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	/* Deklarasi array yang nilainya diset di awal, boleh tidak dituliskan jumlah lebar array-nya, cukup ganti dengan tanda 3 titik (...).
	Metode penulisan ini membuat kapasitas array otomatis dihitung dari jumlah elemen array yang ditulis. */
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numbers)

	// Perulangan Elemen Array Menggunakan Keyword for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	//Perulangan Elemen Array Menggunakan Keyword for - range
	for _, buah := range fruits {
		fmt.Printf("%s\n", buah)
	}
}
