package controller

import (
	"node/database"
	"node/model"
)

var KomikList []database.Komik
var komikIDCounter int = 1

func TambahKomik(
	judul, penulis, ilustrator, genre string,
	tahunTerbit int,
	penerbit, sinopsis string,
	jumlahHalaman int,
	status string,
	rating float64,
) {
	komik := model.CreateKomik(
		komikIDCounter, judul, penulis, ilustrator, genre,
		tahunTerbit, penerbit, sinopsis, jumlahHalaman, status, rating,
	)
	KomikList = append(KomikList, komik)
	komikIDCounter++
}
