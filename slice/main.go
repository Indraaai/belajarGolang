package main

import "fmt"

func main() {

	fmt.Println("hello slice")

	// Inisialisasi Slice
	/* slice bentuknya sama seperti array namun tidak perlu mendefiniskan jumlah elemen diawal
	Salah satu perbedaan slice dan array bisa diketahui pada saat deklarasi variabel-nya,
	jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah slice. */

	var fruits = []string{"indra", "bayu", "muji"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

}
