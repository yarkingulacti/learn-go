package main

import "fmt"

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
}

func sliceTopla[T int | float64](dilim []T) T {
	var toplam T
	for _, deger := range dilim {
		toplam += deger
	}
	return toplam
}
