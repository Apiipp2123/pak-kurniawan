package komik

import (
    "strings"
)

func TambahKomik(k Komik) {
    k.ID = len(Koleksi) + 1
    Koleksi = append(Koleksi, k)
}

func HapusKomik(id int) bool {
    for i, k := range Koleksi {
        if k.ID == id {
            Koleksi = append(Koleksi[:i], Koleksi[i+1:]...)
            return true
        }
    }
    return false
}

func UpdateKomik(id int, updated Komik) bool {
    for i, k := range Koleksi {
        if k.ID == id {
            updated.ID = id
            Koleksi[i] = updated
            return true
        }
    }
    return false
}

func CariKomik(keyword string) []Komik {
    var hasil []Komik
    keyword = strings.ToLower(keyword)

    for _, k := range Koleksi {
        if strings.Contains(strings.ToLower(k.Judul), keyword) ||
            strings.Contains(strings.ToLower(k.Genre), keyword) {
            hasil = append(hasil, k)
        }
    }
    return hasil
}
