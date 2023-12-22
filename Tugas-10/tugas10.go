package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"time"
	"flag"
)

// soal 1
func printKalimatDanTahun(sentence string, year int) {
	fmt.Printf("%s %d\n", sentence, year)
}
func yangPentingJalan(sentence string, year int) {
	defer printKalimatDanTahun(sentence, year)
	fmt.Println("coba-coba belajar goang di sanbercode")
}

// soal 2
func kelilingSegitigaSamaSisi(sisi int, keterangan bool) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	if sisi == 0 {
		errMsg := "Maaf anda belum menginput sisi dari segitiga sama sisi"
		if keterangan {
			panic(errMsg)
		}
		return "", errors.New(errMsg)
	}
	keliling := sisi * 3
	if keterangan {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}
	return fmt.Sprintf("%d", keliling), nil
}

// soal 3
func tambahAngka(nilai int, total *int) {
	*total += nilai
}
func cetakAngka(nilai *int) {
	fmt.Printf("%+v\n", *nilai)
}

// soal 4
var phone = []string{}
func tambahMerek(phones *[]string, merek string) {
	*phones = append(*phones, merek)
}
func tampilkanData(phones *[]string) {
	sort.Strings(*phones)
	for i, merek := range *phones {
		fmt.Printf("%d. %s\n", i+1, merek)
		time.Sleep(time.Second * 1)		
	}
}

// soal 5
type bangunDatar interface {
	luas() 
	keliling()
}
type lingkaran struct {
    jariJari float64
}
func (l lingkaran) luas() int {
	luas := math.Pi * math.Pow(l.jariJari, 2)
	return int(math.Round(luas))
}
func (l lingkaran) keliling() int {
	keliling := math.Pi * l.jariJari * 2
	return int(math.Round(keliling))
}

func main() {
	angka := 1

	// soal 1
	kalimat := "Golang Backend Development"
	tahun := 2021
	yangPentingJalan(kalimat, tahun)

	// soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// soal 3
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// soal 4
	var phones = []string{}
	tambahMerek(&phones, "Xiaomi")
	tambahMerek(&phones, "Asus")
	tambahMerek(&phones, "IPhone")
	tambahMerek(&phones, "Samsung")
	tambahMerek(&phones, "Oppo")
	tambahMerek(&phones, "Realme")
	tambahMerek(&phones, "Vivo")
	tampilkanData(&phones)

	// soal 5
	beberapaJariJari := []float64{5,7,15}
	for _, radius := range beberapaJariJari {
		l := lingkaran{jariJari: radius}

		fmt.Printf("Lingkaran dengan jari-jari %.2f:\n", radius)
		fmt.Printf("Luas: %d\n", l.luas())
		fmt.Printf("Keliling: %d\n\n", l.keliling())
	}

	// soal 6
	var panjang int
	var lebar int

	flag.IntVar(&panjang, "panjang", 0, "Panjang dari persegi panjang")
	flag.IntVar(&lebar, "lebar", 0, "Lebar dari persegi panjang")

	flag.Parse()

	if panjang == 0 || lebar == 0 {
        fmt.Println("Panjang dan lebar harus lebih dari 0")
        return
    }
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Printf("Luas persegi panjang: %d\n", luas)
	fmt.Printf("Keliling persegi panjang: %d\n", keliling)
}