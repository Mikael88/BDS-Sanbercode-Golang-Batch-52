package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mahasiswa struct {
	gorm.Model
	Nama        string
	MataKuliah  string
	IndeksNilai string
	Nilai       uint
}

var db *gorm.DB

func init() {
	dsn := "root:Mikanusan8898!@tcp(127.0.0.1:3306)/tugas14?parseTime=true"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Mahasiswa{})
}

func handleMahasiswa(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMahasiswa(w, r)
	case http.MethodPost:
		createMahasiswa(w, r)
	case http.MethodPut:
		updateMahasiswa(w, r)
	case http.MethodDelete:
		deleteMahasiswa(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	var mahasiswas []Mahasiswa
	db.Find(&mahasiswas)

	// Mengirim data sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mahasiswas)
}

func createMahasiswa(w http.ResponseWriter, r *http.Request) {
	var mahasiswa Mahasiswa
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	mahasiswa.Nama = r.FormValue("Nama")
	mahasiswa.MataKuliah = r.FormValue("MataKuliah")
	mahasiswa.Nilai = parseUint(r.FormValue("Nilai"))

	// Validasi nilai
	if mahasiswa.Nilai > 100 {
		http.Error(w, "Nilai maksimal adalah 100", http.StatusBadRequest)
		return
	}

	// Mengisi IndeksNilai sesuai kondisi
	switch {
	case mahasiswa.Nilai >= 80:
		mahasiswa.IndeksNilai = "A"
	case mahasiswa.Nilai >= 70:
		mahasiswa.IndeksNilai = "B"
	case mahasiswa.Nilai >= 60:
		mahasiswa.IndeksNilai = "C"
	case mahasiswa.Nilai >= 50:
		mahasiswa.IndeksNilai = "D"
	default:
		mahasiswa.IndeksNilai = "E"
	}

	// Menambahkan data ke database
	db.Create(&mahasiswa)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Data Mahasiswa berhasil ditambahkan")
}

func updateMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID dari URL
	// Misalnya: /mahasiswa/1
	id := r.URL.Path[len("/mahasiswa/"):]

	// Mencari mahasiswa berdasarkan ID
	var mahasiswa Mahasiswa
	result := db.First(&mahasiswa, id)
	if result.Error != nil {
		http.Error(w, "Mahasiswa tidak ditemukan", http.StatusNotFound)
		return
	}

	// Mendapatkan data dari form
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Memperbarui data mahasiswa
	db.Model(&mahasiswa).Updates(Mahasiswa{
		Nama:        r.FormValue("Nama"),
		MataKuliah:  r.FormValue("MataKuliah"),
		Nilai:       parseUint(r.FormValue("Nilai")),
		IndeksNilai: calculateIndeksNilai(parseUint(r.FormValue("Nilai"))),
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Data Mahasiswa berhasil diperbarui")
}

func deleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID dari URL
	// Misalnya: /mahasiswa/1
	id := r.URL.Path[len("/mahasiswa/"):]

	// Mencari mahasiswa berdasarkan ID
	var mahasiswa Mahasiswa
	result := db.First(&mahasiswa, id)
	if result.Error != nil {
		http.Error(w, "Mahasiswa tidak ditemukan", http.StatusNotFound)
		return
	}

	// Menghapus data mahasiswa
	db.Delete(&mahasiswa)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Data Mahasiswa berhasil dihapus")
}

func parseUint(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 32)
	return uint(v)
}

func calculateIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}