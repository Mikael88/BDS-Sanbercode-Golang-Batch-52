package main

import "fmt"

// soal 1
type buah struct{
	Nama, Warna string
	AdaBijinga bool
	Harga int
}
func cetakInfoBuahA(b buah) {
    fmt.Printf("%s      ",b.Nama)
	fmt.Printf("%s         ",b.Warna)
	if b.AdaBijinga {
		fmt.Printf("%s              ", "Ada")
	} else {
		fmt.Printf("%s        ", "Tidak Ada")
	}
	fmt.Printf("%d\n", b.Harga)
}
func cetakInfoBuahB(b buah) {
    fmt.Printf("%s   ",b.Nama)
	fmt.Printf("%s  ",b.Warna)
	if b.AdaBijinga {
		fmt.Printf("%s              ", "Ada")
	} else {
		fmt.Printf("%s        ", "Tidak Ada")
	}
	fmt.Printf("%d\n", b.Harga)
}
func cetakInfoBuahC(b buah) {
    fmt.Printf("%s     ",b.Nama)
	fmt.Printf("%s         ",b.Warna)
	if b.AdaBijinga {
		fmt.Printf("%s              ", "Ada")
	} else {
		fmt.Printf("%s        ", "Tidak Ada")
	}
	fmt.Printf("%d\n", b.Harga)
}

// soal 2
type segitiga struct{
	alas, tinggi int
  }
  func (segitiga segitiga) hitungLuasSegitiga() {
	fmt.Printf("Luas segitiga adalah %d\n", (segitiga.alas * segitiga.tinggi) / 2)
}

  type persegi struct{
	sisi int
  }
  func (persegi persegi) hitungLuasPersegi() {
	fmt.Printf("Luas persegi adalah %d\n", persegi.sisi * persegi.sisi)
  }

  type persegiPanjang struct{
	panjang, lebar int
  }
  func (persegiPanjang persegiPanjang) hitungLuasPersegiPanjang() {
    fmt.Printf("Luas persegi panjang adalah %d\n", persegiPanjang.panjang * persegiPanjang.lebar)
  }

// soal 3
type phone struct{
	name, brand string
	year int
	colors []string
 }
func (p *phone)tambahWarna(warnaBaru string) {
	p.colors = append(p.colors, warnaBaru)
}
func cetakPhoneInfo(myPhone *phone) {
	fmt.Printf("%s\n", myPhone.name)
    fmt.Printf("%s\n", myPhone.brand)
    fmt.Printf("%d\n", myPhone.year)
    for i := 0; i < len(myPhone.colors); i++ {
        fmt.Printf("%s ", myPhone.colors[i])
    }
	fmt.Println()
}

// soal 4
type movie struct{
	title, genre string
	duration, year int
}
func (m *movie) tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	newMovie := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}

	*dataFilm = append(*dataFilm, newMovie)
}

func main() {
// soal 1
var nanas = buah{"Nanas", "Kuning", false, 9000}
var jeruk = buah{"Jeruk", "Oranye", true, 8000}
var semangka = buah{"Semangka", "Hijau & Merah", true, 10000}
var pisang = buah{"Pisang", "Kuning", false, 5000}
fmt.Print("Nama       Warna          Ada Bijinya      Harga\n")
cetakInfoBuahA(nanas)
cetakInfoBuahA(jeruk)
cetakInfoBuahB(semangka)
cetakInfoBuahC(pisang)

// soal 2
segitiga1 := segitiga{5, 8}
segitiga1.hitungLuasSegitiga()
persegi1 :=persegi{5}
persegi1.hitungLuasPersegi()
persegiPanjang1 := persegiPanjang{5, 8}
persegiPanjang1.hitungLuasPersegiPanjang()

// soal 3
myPhone := phone{
	name: "Galaxy S25",
    brand: "Samsung",
    year: 2024,
    colors: []string{
		"Red", 
		"Black",
	},
}
 cetakPhoneInfo(&myPhone)
 myPhone.tambahWarna("White")
 cetakPhoneInfo(&myPhone)

 // soal 4
 var dataFilm = []movie{}
m := movie{}
m.tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
m.tambahDataFilm("Avenger", 120, "action", 2019, &dataFilm)
m.tambahDataFilm("Spiderman", 120, "action", 2004, &dataFilm)
m.tambahDataFilm("Juon", 120, "horror", 2004, &dataFilm)

fmt.Println("Data Film:")
for i, film := range dataFilm {
	fmt.Printf("%d. Title: %s\n   Duration: %d minutes\n   Genre: %s\n   Year: %d\n", i+1, film.title, film.duration, film.genre, film.year)
}
}