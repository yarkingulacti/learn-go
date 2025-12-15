package main

import "fmt"

type kus struct {
	cins    string
	renk    string
	habitat string
	yas     int
}

type kopek struct {
	cins     string
	renk     string
	yas      string
	safkanMi bool
}

type hayvan[T kus | kopek] struct {
	bilgi T
	isim  string
}

/*
	Go'da Generics (Tür Parametreleri):
		- Fonksiyonlar ve veri yapıları için tür parametreleri tanımlamaya olanak tanır
		- Tekrarlayan kodu önler ve tür güvenliğini artırır
		- 'type' anahtar kelimesi ile tür parametreleri belirtilir
		- Go 1.18 ve sonrası sürümlerde kullanılabilir
*/
func turParametreleri() {
	intDilimi := []int{1, 2, 3, 4, 5}
	floatDilimi := []float64{1.5, 2.5, 3.5}
	fmt.Println("🔢 int dilimi toplamı:", sliceTopla(intDilimi))
	fmt.Println("🔢 float dilimi toplamı:", sliceTopla(floatDilimi))

	kusum := hayvan[kus]{bilgi: kus{cins: "Serçe", renk: "Kahverengi", habitat: "Orman", yas: 2}, isim: "Cici"}
	kopegim := hayvan[kopek]{bilgi: kopek{cins: "Golden Retriever", renk: "Altın", yas: "3", safkanMi: true}, isim: "Gofret"}

	fmt.Printf("🐦 Hayvan türü: %v, Cinsi: %v, İsmi: %v\n", "Kuş", kusum.bilgi.cins, kusum.isim)
	fmt.Printf("🐶 Hayvan türü: %v, Cinsi: %v, İsmi: %v\n", "Köpek", kopegim.bilgi.cins, kopegim.isim)
}

func sliceTopla[T int | float64](dilim []T) T {
	var toplam T
	for _, deger := range dilim {
		toplam += deger
	}
	return toplam
}
