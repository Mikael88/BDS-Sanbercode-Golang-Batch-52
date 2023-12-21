package main

import (
    "fmt"
    "Tugas-9/library"
)
func main() {
	segitiga := library.SegitigaSamaSisi{Alas: 5, Tinggi: 8}
	fmt.Println("Luas Segitiga:", segitiga.Luas())
	fmt.Println("Keliling Segitiga:", segitiga.Keliling())

	persegi := library.PersegiPanjang{Panjang: 10, Lebar: 10}
	fmt.Println("Luas Persegi:", persegi.Luas())
	fmt.Println("Keliling Persegi:", persegi.Keliling())

	tabung := library.Tabung{JariJari: 10, Tinggi: 10}
	fmt.Println("Luas Tabung:", tabung.LuasPermukaan())
	fmt.Println("Volume Tabung:", tabung.Volume())

	balok := library.Balok{Panjang: 10, Lebar: 10, Tinggi: 10}
	fmt.Println("Luas Balok:", balok.LuasPermukaan())
	fmt.Println("Volume Balok:", balok.Volume())

	myPhone := library.Phone{
		Name:   "xiaomi note 9",
		Brand:  "xiaomi",
		Year:   2020,
		Colors: []string{"Green", "Yellow", "Blue"},
	}
	fmt.Println(myPhone.ShowPhone())

	fmt.Println(library.LuasPersegi(4, true))
	fmt.Println(library.LuasPersegi(8, false))
	fmt.Println(library.LuasPersegi(0, true))
	fmt.Println(library.LuasPersegi(0, false))
}