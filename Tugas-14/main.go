package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/mahasiswa", handleMahasiswa)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}