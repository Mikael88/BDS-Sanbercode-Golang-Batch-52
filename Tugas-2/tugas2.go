package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
// soal 1
var a,b,c,d,e string = "Bootcamp", "Digital", "Skill", "Sanbercode", "Golang"
fmt.Println(a,b,c,d,e)

// soal 2
halo := "Halo Dunia"
find := "Dunia"
replaceWith := "Golang"

newText := strings.Replace(halo, find, replaceWith, 1)
fmt.Println(newText)

// soal 3
var kataPertama = "saya";
var kataKedua = "senang";
var kataKetiga = "belajar";
var kataKeempat = "golang";

//saya Senang belajaR GOLANG
newKataKedua := strings.Title(kataKedua)
last := kataKetiga[len(kataKetiga)-1]
upped := strings.ToUpper(string(last))
newKataKetiga := kataKetiga[:len(kataKetiga)-1] + upped
newKataKeempat := strings.ToUpper(kataKeempat)

fmt.Println(kataPertama, newKataKedua, newKataKetiga, newKataKeempat)

// soal 4
var angkaPertama= "8";
var angkaKedua= "5";
var angkaKetiga= "6";
var angkaKeempat = "7";

var newAngkaPertama,_ = strconv.Atoi(angkaPertama)
var newAngkaKedua,_ = strconv.Atoi(angkaKedua)
var newAngkaKetiga,_ = strconv.Atoi(angkaKetiga)
var newAngkaKeempat,_ = strconv.Atoi(angkaKeempat)

var result = newAngkaPertama + newAngkaKedua + newAngkaKetiga + newAngkaKeempat

fmt.Println(result)

// soal 5
kalimat := "halo halo bandung"
angka := 2021

findHalo := "halo"
replaceHi := "Hai"

newKalimat := strings.Replace(kalimat, findHalo, replaceHi, 2)

fmt.Printf("\"%s\" - %d\n", newKalimat, angka)
}