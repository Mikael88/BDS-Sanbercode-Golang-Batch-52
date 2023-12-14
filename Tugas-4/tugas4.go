package main

import "fmt"

func main() {
	// soal 1
	for i := 0; i <= 10; i++ {
		if i%3 == 0 && i%2 == 1{
			fmt.Println(i, "- I Love Coding")
		} else if i%2 == 0 {
			fmt.Println(i, "- Santai")
		} else if i%2 == 1 {
			fmt.Println(i, "- Berkualitas")
		} 
	}

	// soal 2
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		} 
		fmt.Println()
	}

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	slice := kalimat[2:]
	fmt.Println(slice)

	// soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam")
	sayuran = append(sayuran, "Buncis")
	sayuran = append(sayuran, "Kangkung")
	sayuran = append(sayuran, "Kubis")
	sayuran = append(sayuran, "Seledri")
	sayuran = append(sayuran, "Tauge")
	sayuran = append(sayuran, "Timun")

	for index, sayur := range sayuran {
		fmt.Printf("%d. %s\n", index+1, sayur)
	}

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	  }
	var volumeBalok = satuan["panjang"]*satuan["lebar"]*satuan["tinggi"]
	for keterangan, besaran := range satuan {
		fmt.Printf("%s = %d\n", keterangan, besaran)
	}
	fmt.Printf("Volume balok = %d", volumeBalok)
}