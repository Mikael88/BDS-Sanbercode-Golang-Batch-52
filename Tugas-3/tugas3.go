package main

import (
	"fmt"
	"strconv"
)

// soal 1
var panjangPersegiPanjang string = "8"
var lebarPersegiPanjang string = "5"
var alasSegitiga string = "6"
var tinggiSegitiga string = "7"

var newPanjangPersegiPanjang,_ = strconv.Atoi(panjangPersegiPanjang)
var newLebarPersegiPanjang,_ = strconv.Atoi(lebarPersegiPanjang)
var newAlasSegitiga,_ = strconv.Atoi(alasSegitiga)
var newTinggiSegitiga,_ = strconv.Atoi(tinggiSegitiga)

var luasPersegiPanjang int = newPanjangPersegiPanjang * newLebarPersegiPanjang 
var kelilingPersegiPanjang int = newLebarPersegiPanjang * 2 + newLebarPersegiPanjang * 2
var luasSegitiga int = newAlasSegitiga * newTinggiSegitiga / 2

// soal 2
var nilaiJohn = 80
var nilaiDoe = 50

// soal 3
var tanggal = 8
var bulan = 8
var tahun = 1998

var newTanggal = strconv.Itoa(tanggal)
var newTahun = strconv.Itoa(tahun)
var newBulan string

// soal 4
var contohTahunLahir int = 1998
var generasi string

func main() {
// soal 1
fmt.Println("Luas persegi panjang:", luasPersegiPanjang)
fmt.Println("Keliling persegi panjang:", kelilingPersegiPanjang)
fmt.Println("Luas segitiga:", luasSegitiga)

// soal 2
if nilaiJohn >= 80 {
	fmt.Println("Nilai John: A")
} else if nilaiJohn >= 70 && nilaiJohn < 80 {
	fmt.Println("Nilai John: B")
} else if nilaiJohn >= 60 && nilaiJohn < 70 {
	fmt.Println("Nilai John: C")
} else if nilaiJohn >= 50 && nilaiJohn < 60 {
	fmt.Println("Nilai John: D")
} else {
	fmt.Println("Nilai John: E")
}

if nilaiDoe >= 80 {
	fmt.Println("Nilai Doe: A")
} else if nilaiDoe >= 70 && nilaiDoe < 80 {
	fmt.Println("Nilai Doe: B")
} else if nilaiDoe >= 60 && nilaiDoe < 70 {
	fmt.Println("Nilai Doe: C")
} else if nilaiDoe >= 50 && nilaiDoe < 60 {
	fmt.Println("Nilai Doe: D")
} else {
	fmt.Println("Nilai Doe: E")
}

// soal 3
switch bulan {
case 1:
	newBulan = "Januari"
case 2:
	newBulan = "Febuari"
case 3:
	newBulan = "Maret"
case 4:
	newBulan = "April"
case 5:
	newBulan = "Mei"
case 6:
	newBulan = "Juni"
case 7:
	newBulan = "Juli"
case 8:
	newBulan = "Agustus"
case 9:
	newBulan = "September"
case 10:
	newBulan = "Oktober"
case 11:
	newBulan = "November"
case 12:
	newBulan = "Desember"
default:
	newBulan = "Bulan tidak ada"
}
result := fmt.Sprintf("%s %s %s", newTanggal, newBulan, newTahun)
fmt.Println("Tanggal lahir:", result)

// soal 4
switch {
case contohTahunLahir >= 1944 && contohTahunLahir <= 1964:
	generasi = "Baby boomer"
case contohTahunLahir >= 1965 && contohTahunLahir <= 1979:
	generasi = "Generasi X"
case contohTahunLahir >= 1980 && contohTahunLahir <= 1994:
	generasi = "Generasi Y (Millennials)"
case contohTahunLahir >= 1995 && contohTahunLahir <= 2015:
	generasi = "Generasi Z"
default:
	generasi = "Tidak dapat dikategorikan dalam generasi yang telah ditentukan"
}
fmt.Printf("Tahun kelahiran: %d, Kategori: %s\n", contohTahunLahir, generasi)

}