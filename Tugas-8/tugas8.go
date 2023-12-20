package main

import (
	"fmt"
	"math"
	"strings"
)

// soal 1
//////////////////////////////////////////
type segitigaSamaSisi struct{
  alas, tinggi int
}
func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}
func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}
//////////////////////////////////////////
type persegiPanjang struct{
  panjang, lebar int
}
func (p persegiPanjang) keliling() int {
    return 2 * (p.panjang + p.lebar)
}
func (p persegiPanjang) luas() int {
    return p.panjang * p.lebar
}
//////////////////////////////////////////
type tabung struct{
  jariJari, tinggi float64
}
func (t tabung) luasPermukaan() float64 {
    return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}
func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}
//////////////////////////////////////////
type balok struct{
  panjang, lebar, tinggi int
}
func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}
func (b balok) volume() float64 {
  return float64(b.panjang * b.lebar * b.tinggi)
}
//////////////////////////////////////////
type hitungBangunDatar interface{
  luas() int
  keliling() int
}
type hitungBangunRuang interface{
  volume() float64
  luasPermukaan() float64
}

// soal 2
type phone struct{
  name, brand string
  year int
  colors []string
}

func (p phone) showPhone() string {
  colorStr := strings.Join(p.colors,", ")
  result := fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolor : %s", p.name, p.brand, p.year, colorStr)
  return result
}

type phoneIdentifier interface{
  showPhone() string
}

// soal 3
func luasPersegi(sisi int, keterangan bool) any {
  if sisi == 0 {
    if keterangan {
      return "Maaf anda belum menginput sisi persegi"
    } else {
      return nil
    }
  }
  luas := sisi * sisi
  if keterangan {
    return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
  }
  return luas
}

func main() {
// soal 1
segitiga := segitigaSamaSisi{5,8}
fmt.Println("Luas Segitiga:", segitiga.luas())
fmt.Println("Keliling Segitiga:", segitiga.keliling())

persegi := persegiPanjang{10,10}
fmt.Println("Luas Persegi:", persegi.luas())
fmt.Println("Keliling Persegi:", persegi.keliling())

tabung := tabung{10,10}
fmt.Println("Luas Tabung:", tabung.luasPermukaan())
fmt.Println("Volume Tabung:", tabung.volume())

balok := balok{10,10,10}
fmt.Println("Luas Balok:", balok.luasPermukaan())
fmt.Println("Volume Balok:", balok.volume())

// soal 2
myPhone := phone{
  name:   "xiaomi note 9",
  brand:  "xiaomi",
  year:   2020,
  colors: []string{"Green", "Yellow", "Blue"},
}
fmt.Println(myPhone.showPhone())

// soal 3
fmt.Println(luasPersegi(4, true))
fmt.Println(luasPersegi(8, false))
fmt.Println(luasPersegi(0, true))
fmt.Println(luasPersegi(0, false))

// soal 4
var prefix interface{}= "hasil penjumlahan dari "
var kumpulanAngkaPertama interface{} = []int{6,8}
var kumpulanAngkaKedua interface{} = []int{12,14}
angkaPertama, ok1 := kumpulanAngkaPertama.([]int)
angkaKedua, ok2 := kumpulanAngkaKedua.([]int)
if ok1 && ok2 {
	total := 0
	for _, angka := range angkaPertama {
		total += angka
	}
	for _, angka := range angkaKedua {
		total += angka
	}
	fmt.Printf("%s", prefix)
	for i, angka := range append(angkaPertama, angkaKedua...) {
		fmt.Printf("%d", angka)
		if i < len(angkaPertama)+len(angkaKedua)-1 {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", total)
} else {
		fmt.Println("Tipe data kumpulan angka tidak sesuai")
}
}