package model

import (
	"node/database"
	"strings"
)

func CreateKomik(
	id int,
	judulKomik, penulisKomik, ilustratorKomik, genreKomik string,
	tahunTerbit int,
	penerbitKomik, sinopsisKomik string,
	jumlahHalaman int,
	statusKomik string,
	ratingKomik float64,
) database.Komik {
	return database.Komik{
		ID:            id,
		Judul:         strings.TrimSpace(judulKomik),
		Penulis:       strings.TrimSpace(penulisKomik),
		Ilustrator:    strings.TrimSpace(ilustratorKomik),
		Genre:         strings.TrimSpace(genreKomik),
		TahunTerbit:   tahunTerbit,
		Penerbit:      strings.TrimSpace(penerbitKomik),
		Sinopsis:      strings.TrimSpace(sinopsisKomik),
		JumlahHalaman: jumlahHalaman,
		Status:        strings.TrimSpace(statusKomik),
		Rating:        ratingKomik,
	}
}
