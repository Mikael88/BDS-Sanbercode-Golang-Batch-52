package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

// soal 2
var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

func getMovies(ch chan string, movies ...string) {
	defer close(ch)
	for i, movie := range movies {
		time.Sleep(100 * time.Millisecond)
        ch <- fmt.Sprintf("%d. %s", i+1, movie)
    }
}

// soal 3
type Result struct {
    jariJari          float64
    luasLingkaran     float64
    kelilingLingkaran float64
    volumeTabung      float64
}
func perhitungan(r float64, ch chan Result, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	result2 := Result{
		jariJari: r,
		luasLingkaran: r * r * math.Pi,
        kelilingLingkaran: 2 * r * math.Pi,
        volumeTabung: r * r * r * math.Pi,
	}
	ch <- result2
}

// soal 4
type hasilPerhitungan struct {
	Panjang  int
	Lebar    int
	Tinggi   int
	Luas     int
	Keliling int
	Volume   int
	jenisPerhitungan string
}
func luasPersegiPanjang(panjang, lebar int, ch3 chan hasilPerhitungan, wg3 *sync.WaitGroup) {
	defer wg3.Done()
	result := hasilPerhitungan{
		Panjang: panjang,
		Lebar: lebar,
        Tinggi: panjang + lebar,
        Luas: panjang * lebar,
        jenisPerhitungan: "Luas",
	}
	ch3<- result
}
func kelilingPersegiPanjang(panjang, lebar int, ch3 chan hasilPerhitungan, wg3 *sync.WaitGroup) {
    defer wg3.Done()
    result := hasilPerhitungan{
        Panjang: panjang,
		Lebar: lebar,
		Keliling: 2 * (panjang + lebar),
        jenisPerhitungan: "Keliling",
    }
    ch3 <- result
}
func volumeTabung(panjang, lebar, tinggi int, ch3 chan hasilPerhitungan, wg3 *sync.WaitGroup) {
	defer wg3.Done()
	result := hasilPerhitungan{
		Panjang: panjang,
        Lebar: lebar,
        Tinggi: tinggi,
        Volume: panjang * lebar * tinggi,
        jenisPerhitungan: "Volume",
	}
	ch3 <- result	
}

func main() {
	// soal 1
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	var wg sync.WaitGroup
	resultCh := make(chan string, len(phones))

	wg.Add(1)
	go func() {
		defer wg.Done()
        sort.Strings(phones)
    }()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		for i, phone := range phones {
            resultCh <- fmt.Sprintf("%d. %s", i+1, phone)
			time.Sleep(100 * time.Millisecond)
        }
		close(resultCh)
	}()
	
	printWG := sync.WaitGroup{}
	printWG.Add(1)
	go func ()  {
		defer printWG.Done()
        for result := range resultCh {
            fmt.Println(result)
        }
	}()

	wg.Wait()
	printWG.Wait()

	// soal 2
	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	fmt.Println("List Movies:")
	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	radii := []float64{8, 14, 20}
    var wg2 sync.WaitGroup
    resultCh2 := make(chan Result, len(radii))

    for _, r := range radii {
        wg2.Add(1)
        go perhitungan(r, resultCh2, &wg2)
    }

    go func() {
        wg2.Wait()
        close(resultCh2)
    }()

    for result2 := range resultCh2 {
        fmt.Printf("Radius: %.2f\n", result2.jariJari)
        fmt.Printf("Circle Area: %.2f\n", result2.luasLingkaran)
        fmt.Printf("Circle Circumference: %.2f\n", result2.kelilingLingkaran)
        fmt.Printf("Cylinder Volume: %.2f\n", result2.volumeTabung)
        fmt.Println("-----")
    }

	// soal 4
	panjang := 5
	lebar := 3
	tinggi := 2

	var wg3 sync.WaitGroup
	resultCh3 := make(chan hasilPerhitungan, 3)
	
	wg3.Add(3)
	go luasPersegiPanjang(panjang, lebar, resultCh3, &wg3)
	go kelilingPersegiPanjang(panjang, lebar, resultCh3, &wg3)
	go volumeTabung(panjang, lebar, tinggi, resultCh3, &wg3)

	go func() {
        wg3.Wait()
        close(resultCh3)
    }()
	for {
        select {
        case result, ok := <-resultCh3:
            if !ok {
                return
            }
            switch result.jenisPerhitungan {
            case "Luas":
                fmt.Printf("Luas Persegi Panjang: %d\n", result.Luas)
            case "Keliling":
                fmt.Printf("Keliling Persegi Panjang: %d\n", result.Keliling)
            case "Volume":
                fmt.Printf("Volume Tabung: %d\n", result.Volume)
            }
        }
    }
}