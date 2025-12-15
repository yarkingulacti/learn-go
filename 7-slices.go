package main

import "fmt"

/*
	Go'da Dilimler (Slices):
		- Dinamik, dizinin elemanlarına esnek bir görünüm
		- Dilim literal veya make fonksiyonu ile oluşturulur
		- Uzunluk (length) ve kapasite (capacity) özelliklerine sahiptir
*/
func dilimler() {
	// Dilim literal ile oluşturma, boyutsuz dizi literaline benzer
	intDilimi := []int32{1, 2, 3, 4, 5}
	fmt.Println("📦 Dilim örneği:", intDilimi)
	fmt.Println("↔️ Uzunluk:", len(intDilimi)) // dilimin uzunluğu
	fmt.Println("📐 Kapasite:", cap(intDilimi)) // dilimin kapasitesi

	intDilimi = append(intDilimi, 6)
	fmt.Println("➕ Append sonrası dilim:", intDilimi)
	fmt.Println("↔️ Yeni uzunluk:", len(intDilimi))
	fmt.Println("📐 Yeni kapasite:", cap(intDilimi))

	intDilimi2 := []int32{7, 8, 9}
	intDilimi = append(intDilimi, intDilimi2...) // başka bir dilimi eklerken ... kullan

	fmt.Println("📦 Başka bir dilim eklenmiş hali:", intDilimi)

	uzunluk := 3
	kapasite := 5

	// make ile belirtilen uzunluk ve kapasite ile dilim oluşturma
	intDilimi3 := make([]int32, uzunluk, kapasite)
	fmt.Println("🛠️ make ile oluşturulan dilim (başlangıç):", intDilimi3)
	fmt.Println("↔️ Uzunluk:", len(intDilimi3))
	fmt.Println("📐 Kapasite:", cap(intDilimi3))
}
