package main

import "fmt"

/*
	Go'da Diziler:
		- Aynı türde sabit boyutlu eleman koleksiyonu
		- Belirli bir boyut ile tanımlanır
		- Elemanlara indeksler aracılığıyla erişilir
		- Varsayılan olarak sıfır değerine sahiptir (ör. int için 0, string için "")
*/
func diziler() {
	var intDizi [5]int8 = [5]int8{1, 2, 3, 4, 5}
	fmt.Println("🧩 Dizi örneği:", intDizi)
	fmt.Println("🔢 İlk eleman:", intDizi[0])
	fmt.Println("🔪 Dilimleme (1:3):", intDizi[1:3])
	intDizi[4] = 10
	fmt.Println("✏️ Değiştirilmiş 4. indeks:", intDizi[4])

	// dizi adresleri örneği
	fmt.Println("📍 Dizi elemanlarının bellek adresleri:", &intDizi[0], &intDizi[1], &intDizi[2])
}
