package main

import "fmt"

// Penggunaan Map
/*  penulisan map adalah nama varaiabel= map[tipe data key]value{}
tipe data map seperti tipe data objek di javasripst dengan key value
*/

func main() {
	fmt.Println("Inisialisasi Nilai Map")

	var ayam = map[string]int{
		"jantan": 1,
		"betina": 2,
		"waria":  3,
	}

	// Iterasi Item Map Menggunakan for - range
	for key, value := range ayam {
		fmt.Println(key, "  \t:", value)
	}

	//Menghapus Item Map dengan fungsi delete
	// untuk menghgapusmya gunakan varaiabel nya lalu ambil keynya
	fmt.Println("panjang data map adalah =", len(ayam))

	delete(ayam, "waria")

	fmt.Println(ayam)

	// Kombinasi Slice & Map
	/* penerapannay adalah tipe data slice[],map[tipe data mapa]tipe data value*/
	var ayam1 = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	for _, val := range ayam1 {
		fmt.Println(val)
	}
}
