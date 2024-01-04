package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}
var users = map[string]string{
	"garuda": "garuda",
}

var mu sync.Mutex

func main() {
	http.HandleFunc("/nilai-mahasiswa", handleNilaiMahasiswa)
	http.HandleFunc("/get-nilai-mahasiswa", handleGetNilaiMahasiswa)

	// Menjalankan server pada port 8080
	http.ListenAndServe(":8080", nil)
}

func handleNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Middleware untuk Basic Auth
	username, password, ok := basicAuth(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verifikasi username dan password
	if storedPassword, found := users[username]; !found || storedPassword != password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Menerima data dalam bentuk JSON atau formData
	var inputNilaiMahasiswa NilaiMahasiswa
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		inputNilaiMahasiswa.Nama = r.FormValue("Nama")
		inputNilaiMahasiswa.MataKuliah = r.FormValue("MataKuliah")
		inputNilaiMahasiswa.Nilai = parseUint(r.FormValue("Nilai"))
	}

	// Validasi nilai
	if inputNilaiMahasiswa.Nilai > 100 {
		http.Error(w, "Nilai maksimal adalah 100", http.StatusBadRequest)
		return
	}

	// Mengisi IndeksNilai sesuai kondisi
	switch {
	case inputNilaiMahasiswa.Nilai >= 80:
		inputNilaiMahasiswa.IndeksNilai = "A"
	case inputNilaiMahasiswa.Nilai >= 70:
		inputNilaiMahasiswa.IndeksNilai = "B"
	case inputNilaiMahasiswa.Nilai >= 60:
		inputNilaiMahasiswa.IndeksNilai = "C"
	case inputNilaiMahasiswa.Nilai >= 50:
		inputNilaiMahasiswa.IndeksNilai = "D"
	default:
		inputNilaiMahasiswa.IndeksNilai = "E"
	}

	// Mengolah ID
	mu.Lock()
	defer mu.Unlock()
	inputNilaiMahasiswa.ID = uint(len(nilaiNilaiMahasiswa) + 1)

	// Menambahkan data ke nilaiNilaiMahasiswa
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, inputNilaiMahasiswa)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Data Nilai Mahasiswa berhasil ditambahkan")
}

func handleGetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Mengirim data sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func basicAuth(r *http.Request) (username, password string, ok bool) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", "", false
	}

	decodedAuth, err := base64.StdEncoding.DecodeString(authHeader[len("Basic "):])
	if err != nil {
		return "", "", false
	}

	credentials := string(decodedAuth)
	usernamePassword := strings.Split(credentials, ":")
	if len(usernamePassword) != 2 {
		return "", "", false
	}

	return usernamePassword[0], usernamePassword[1], true
}

func parseUint(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 32)
	return uint(v)
}
