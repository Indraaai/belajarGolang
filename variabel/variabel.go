package main

import "fmt"

func main() {
	fmt.Println("hello guys")
	// variabel deklarasi langsung dengan titik 2 diawal setelah nama varaibel dan sebelum =
	name := "indra"
	var lastName string = "firmansyah"
	fmt.Printf("halo %s %s!\n", name, lastName)

	// variabel denga tanda undeskor bisa tidak digunakan istilahnya adalah tempat sampah
	alamat, _ := "blangpulo", "bukitindah"
	fmt.Println(alamat)
}
