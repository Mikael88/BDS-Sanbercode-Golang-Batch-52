package main

import (
	"fmt"
	"strings"
)

// soal 1
func luasPersegiPanjang(p, l int) int {
	return p*l
}
func kelilingPersegiPanjang(p, l int) int {
	return p*2 + l*2
}
func volumeBalok(p, l, t int) int {
	return p*l*t
}

// soal 2
func introduce(nama, gender, pekerjaan, umur string) (result string) {
	var title string
	if gender == "laki-laki" {
		title = "Pak"
	} else if gender == "perempuan" {
		title = "Bu"
	}
	result = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, nama, pekerjaan, umur)
	return result
}

// soal 3
func buahFavorit(nama string, buah ...string) string {
	buahString := strings.Join(buah, "\", \"")
	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah \"%s\"", nama, buahString)
}

// soal 4

func main() {
// soal 1
panjang := 12
lebar := 4
tinggi := 8

luas := luasPersegiPanjang(panjang, lebar)
keliling := kelilingPersegiPanjang(panjang, lebar)
volume := volumeBalok(panjang, lebar, tinggi)

fmt.Println(luas) 
fmt.Println(keliling)
fmt.Println(volume)

// soal 2
john := introduce("John", "laki-laki", "penulis", "30")
fmt.Println(john)
sarah := introduce("Sarah", "perempuan", "model", "28")
fmt.Println(sarah)

// soal 3
buah := []string{"semangka", "jeruk", "melon", "pepaya"}
buahFavoritJohn := buahFavorit("John", buah...)
fmt.Println(buahFavoritJohn)

// soal 4

var dataFilm = []map[string]string{}
// buatlah closure function disini
var tambahDataFilm = func (judul, durasi, genre, tahun string)  {
	film := map[string]string {
		"judul": judul,
		"durasi": durasi,
		"genre": genre,
		"tahun": tahun,
	}
	dataFilm = append(dataFilm, film)
}
tambahDataFilm("LOTR", "2 jam", "action", "1999")
tambahDataFilm("avenger", "2 jam", "action", "2019")
tambahDataFilm("spiderman", "2 jam", "action", "2004")
tambahDataFilm("juon", "2 jam", "horror", "2004")

for _, item := range dataFilm {
   fmt.Println(item)
}
}