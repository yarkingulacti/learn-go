package main

import "fmt"

/*
	For Döngüleri (For Loops) Go'da:
		- Koleksiyonlar üzerinde yineleme yapmak veya kodu tekrarlı çalıştırmak için kullanılır
		- Üç formu vardır: klasik, range tabanlı ve while benzeri
*/
func dongu() {
	intDizi := [5]int{10, 20, 30, 40, 50}
	for indeks, deger := range intDizi {
		fmt.Printf("Index: %v, Value: %v\n", indeks, deger)
	}

	var sayac = 10

	for sayac >= 0 {
		if sayac == 0 {
			break
		}

		fmt.Println(sayac)
		sayac--
	}

	for sayac2 := 0; sayac2 < 5; sayac2++ {
		fmt.Println(sayac2)
	}
}
