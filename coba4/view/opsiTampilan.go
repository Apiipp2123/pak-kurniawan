package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func OpsiTampilan() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== Aplikasi Data Komik CLI ===")
		fmt.Println("1. Tambah Komik")
		fmt.Println("2. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			Inputkomik()
		case "2":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}

		fmt.Println()
	}
}