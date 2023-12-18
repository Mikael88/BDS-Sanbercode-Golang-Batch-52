package main

import "fmt"

// soal 1
func rumusLingkaran(luas *float64, keliling *float64, r float64) {
    *luas = 3.14 * r * r
    *keliling = 3.14 * 2 * r
}

// soal 2
func introduce(kalimat *string, nama, gender, profesi, umur string) {
	var title string
	if gender == "laki-laki" {
		title = "Pak"
	} else if gender == "perempuan" {
		title = "Bu"
	}
	*kalimat = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, nama, profesi, umur)
}

// soal 3
func ubahBuah(buah *[]string, listBuah []string) {
	*buah = listBuah
}
func tampilkanBuah(buah []string) {
	for i, namaBuah := range buah {
		fmt.Printf("%d. %s\n", i+1, namaBuah)
	}
}

// soal 4
func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilm = append(*dataFilm, film)
}

func tampilkanDataFilm(dataFilm *[]map[string]string) {
	for i, film := range *dataFilm {
		fmt.Printf("%d. title: %s\n", i+1, film["title"])
		fmt.Printf("   duration: %s\n", film["duration"])
		fmt.Printf("   genre: %s\n", film["genre"])
		fmt.Printf("   year: %s\n\n", film["year"])
	}
}

func main() {
	// soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64
	rumusLingkaran(&luasLingkaran, &kelilingLingkaran, 7)
	fmt.Println(luasLingkaran)
	fmt.Println(kelilingLingkaran)

	// soal 2
	var sentence string 
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	// soal 3
	var buah = []string{}
	var listBuah = []string{"Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat"}
	ubahBuah(&buah, listBuah)
	tampilkanBuah(buah)

	// soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	tampilkanDataFilm(&dataFilm)
} 	