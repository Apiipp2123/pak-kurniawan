package view

import (
	"bufio"
	"fmt"
	controller "node/controllers"
	"os"
	"strconv"
	"strings"
)

func Inputkomik() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Judul: ")
	judulKomik, _ := reader.ReadString('\n')
	fmt.Print("Penulis: ")
	penulisKomik, _ := reader.ReadString('\n')
	fmt.Print("Ilustrator: ")
	ilustratorKomik, _ := reader.ReadString('\n')
	fmt.Print("Genre: ")
	genreKomik, _ := reader.ReadString('\n')
	fmt.Print("Tahun Terbit: ")
	tahunTerbitStr, _ := reader.ReadString('\n')
	fmt.Print("Penerbit: ")
	penerbitKomik, _ := reader.ReadString('\n')
	fmt.Print("Sinopsis: ")
	sinopsisKomik, _ := reader.ReadString('\n')
	fmt.Print("Jumlah Halaman: ")
	jumlahHalamanStr, _ := reader.ReadString('\n')
	fmt.Print("Status: ")
	statusKomik, _ := reader.ReadString('\n')
	fmt.Print("Rating: ")
	ratingStr, _ := reader.ReadString('\n')

	// Konversi ke tipe data yang sesuai
	tahunTerbit, _ := strconv.Atoi(strings.TrimSpace(tahunTerbitStr))
	jumlahHalaman, _ := strconv.Atoi(strings.TrimSpace(jumlahHalamanStr))
	rating, _ := strconv.ParseFloat(strings.TrimSpace(ratingStr), 64)

	// Kirim ke controller
	controller.TambahKomik(
		judulKomik, penulisKomik, ilustratorKomik, genreKomik,
		tahunTerbit, penerbitKomik, sinopsisKomik,
		jumlahHalaman, statusKomik, rating,
	)

	fmt.Println("✅ Komik berhasil ditambahkan!")
}
