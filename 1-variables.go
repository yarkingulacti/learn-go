package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Go'da Temel Veri Tipleri:
- Tam Sayı Tipleri: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
- Ondalık Sayı Tipleri: float32, float64
- Metin Tipi: string
- Rune Tipi: rune (int32 takma adı, Unicode kod noktası temsil eder)
- Boolean Tipi: bool
*/
func degiskenler() {
	var tamsayi int16 = 32767
	fmt.Println(tamsayi)

	var ondalikAzHassas float32 = 12345678.9 // hassasiyet kaybı
	var ondalikHassas float64 = 12345678.9   // daha hassas

	fmt.Println(ondalikAzHassas)
	fmt.Println(ondalikHassas)

	var ondalik32 float32 = 10.1
	var tamsayi32 int32 = 2
	var sonuc float32 = ondalik32 + float32(tamsayi32) // doğru dönüşüm

	fmt.Println(sonuc)

	var benimMetin string = "Hello, \nGo!"

	fmt.Println(benimMetin)
	fmt.Println(len("ğ"))                    // 2 (UTF-8'de 2 byte)
	fmt.Println(utf8.RuneCountInString("ğ")) // 1 (1 rune)

	var benimRune rune = 'a' // ASCII'de 97
	fmt.Println(benimRune)

	var dogruMu bool = true
	fmt.Println(dogruMu)

	var digerBool bool // varsayılan değer false
	fmt.Println(digerBool)

	var benimMetin2 string
	fmt.Println(benimMetin2) // varsayılan boş string

	var benimRune2 rune
	fmt.Println(benimRune2) // varsayılan 0

	var metinim = "text" // türü string olarak çıkarılır

	fmt.Println(metinim)

	kisaDegisken := "short variable declaration" // tür çıkarımı

	fmt.Println(kisaDegisken)

	var sayi1, sayi2, sayi3 int = 1, 2, 3
	fmt.Println(sayi1, sayi2, sayi3)

	const benimSabit = "This is a constant and cannot be changed. Also, it must be initialized at declaration."

	fmt.Println(benimSabit)
}
