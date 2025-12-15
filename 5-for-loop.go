package main

import "fmt"

/*
	For Döngüleri (For Loops) Go'da:
		- Koleksiyonlar üzerinde yineleme yapmak veya kodu tekrarlı çalıştırmak için kullanılır
		- Üç formu vardır: klasik, range tabanlı ve while benzeri
*/
func dongu() {
	intDizi := [5]int{10, 20, 30, 40, 50}
	fmt.Println("🔁 Range ile yineleme örneği:")
	for indeks, deger := range intDizi {
		fmt.Printf("➡️ İndeks: %v, Değer: %v\n", indeks, deger)
	}

	var sayac = 10
	fmt.Println("⏳ Geri sayım örneği (break ile):")
	for sayac >= 0 {
		if sayac == 0 {
			fmt.Println("🏁 Sayaç sıfırlandı, döngüden çıkılıyor")
			break
		}

		fmt.Println("🔢 Sayaç:", sayac)
		sayac--
	}

	fmt.Println("🔂 Klasik for döngüsü örneği:")
	for sayac2 := 0; sayac2 < 5; sayac2++ {
		fmt.Println("indeks:", sayac2)
	}
}
