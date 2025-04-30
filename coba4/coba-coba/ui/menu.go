package ui

import (
	"bufio"
	"fmt"
	"node/komik"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func TampilkanMenu() {
    for {
        fmt.Println("\n--- Menu Komik ---")
        fmt.Println("1. Tambah Komik")
        fmt.Println("2. Tampilkan Semua Komik")
        fmt.Println("3. Hapus Komik")
        fmt.Println("4. Update Komik")
        fmt.Println("5. Cari Komik (judul/genre)")
        fmt.Println("6. Keluar")
        fmt.Print("Pilih menu: ")
        scanner.Scan()
        input := scanner.Text()

        switch input {
        case "1":
            tambah()
        case "2":
            tampil()
        case "3":
            hapus()
        case "4":
            update()
        case "5":
            cari()
        case "6":
            fmt.Println("Terima kasih!")
            return
        default:
            fmt.Println("Pilihan tidak valid.")
        }
    }
}

func tambah() {
    var k komik.Komik

    fmt.Print("Judul: "); scanner.Scan(); k.Judul = scanner.Text()
    fmt.Print("Penulis: "); scanner.Scan(); k.Penulis = scanner.Text()
    fmt.Print("Ilustrator: "); scanner.Scan(); k.Ilustrator = scanner.Text()
    fmt.Print("Genre: "); scanner.Scan(); k.Genre = scanner.Text()
    fmt.Print("Tahun Terbit: "); scanner.Scan(); k.TahunTerbit, _ = strconv.Atoi(scanner.Text())
    fmt.Print("Penerbit: "); scanner.Scan(); k.Penerbit = scanner.Text()
    fmt.Print("Sinopsis: "); scanner.Scan(); k.Sinopsis = scanner.Text()
    fmt.Print("Jumlah Halaman: "); scanner.Scan(); k.JumlahHalaman, _ = strconv.Atoi(scanner.Text())
    fmt.Print("Status: "); scanner.Scan(); k.Status = scanner.Text()
    fmt.Print("Rating: "); scanner.Scan(); k.Rating, _ = strconv.ParseFloat(scanner.Text(), 64)

    komik.TambahKomik(k)
    fmt.Println("✅ Komik berhasil ditambahkan.")
}

func tampil() {
    if len(komik.Koleksi) == 0 {
        fmt.Println("Belum ada komik.")
        return
    }
    for _, k := range komik.Koleksi {
        fmt.Printf("\n[%d] %s\n", k.ID, strings.ToUpper(k.Judul))
        fmt.Printf("Penulis: %s | Ilustrator: %s | Genre: %s\n", k.Penulis, k.Ilustrator, k.Genre)
        fmt.Printf("Tahun: %d | Halaman: %d | Status: %s | Rating: %.1f\n", k.TahunTerbit, k.JumlahHalaman, k.Status, k.Rating)
        fmt.Printf("Sinopsis: %s\n", k.Sinopsis)
    }
}

func hapus() {
    fmt.Print("Masukkan ID komik yang akan dihapus: ")
    scanner.Scan()
    id, _ := strconv.Atoi(scanner.Text())
    if komik.HapusKomik(id) {
        fmt.Println("✅ Komik berhasil dihapus.")
    } else {
        fmt.Println("❌ Komik tidak ditemukan.")
    }
}

func update() {
    fmt.Print("Masukkan ID komik yang akan diupdate: ")
    scanner.Scan()
    id, _ := strconv.Atoi(scanner.Text())

    var k komik.Komik
    fmt.Print("Judul: "); scanner.Scan(); k.Judul = scanner.Text()
    fmt.Print("Penulis: "); scanner.Scan(); k.Penulis = scanner.Text()
    fmt.Print("Ilustrator: "); scanner.Scan(); k.Ilustrator = scanner.Text()
    fmt.Print("Genre: "); scanner.Scan(); k.Genre = scanner.Text()
    fmt.Print("Tahun Terbit: "); scanner.Scan(); k.TahunTerbit, _ = strconv.Atoi(scanner.Text())
    fmt.Print("Penerbit: "); scanner.Scan(); k.Penerbit = scanner.Text()
    fmt.Print("Sinopsis: "); scanner.Scan(); k.Sinopsis = scanner.Text()
    fmt.Print("Jumlah Halaman: "); scanner.Scan(); k.JumlahHalaman, _ = strconv.Atoi(scanner.Text())
    fmt.Print("Status: "); scanner.Scan(); k.Status = scanner.Text()
    fmt.Print("Rating: "); scanner.Scan(); k.Rating, _ = strconv.ParseFloat(scanner.Text(), 64)

    if komik.UpdateKomik(id, k) {
        fmt.Println("✅ Komik berhasil diupdate.")
    } else {
        fmt.Println("❌ Komik tidak ditemukan.")
    }
}

func cari() {
    fmt.Print("Masukkan judul atau genre: ")
    scanner.Scan()
    keyword := scanner.Text()
    hasil := komik.CariKomik(keyword)

    if len(hasil) == 0 {
        fmt.Println("Tidak ditemukan komik yang cocok.")
        return
    }

    for _, k := range hasil {
        fmt.Printf("\n[%d] %s (%s)\n", k.ID, k.Judul, k.Genre)
    }
}
