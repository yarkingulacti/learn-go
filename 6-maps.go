package main

import "fmt"

/*
	Go'da Haritalar (Maps):
		- Anahtar-değer çiftlerinden oluşan sırasız koleksiyon
		- Anahtarlar benzersizdir ve değerlere erişimde kullanılır
		- make veya map literal ile oluşturulur
		- Varsayılan olarak sıfır değere sahiptir (ör. int için 0, string için "")
*/
func haritalar() {
	// Harita oluşturma
	benimHarita := make(map[string]uint8)
	fmt.Println(benimHarita)

	// Atama sırasında anahtar-değer çifti ekleme
	benimHarita2 := map[string]uint8{"Adam": 30, "Eve": 28}
	fmt.Println(benimHarita2["Adam"])
	fmt.Println(benimHarita2["Sarah"]) // varsayılan değer 0

	yas, varMi := benimHarita2["Sarah"]

	if varMi {
		fmt.Println(yas)
	} else {
		fmt.Println("Key does not exist")
	}

	// bir anahtar-değer çiftini silme
	delete(benimHarita2, "Adam")

	fmt.Println(benimHarita2)

	for anahtar, deger := range benimHarita2 {
		fmt.Printf("Name: %v, Age: %v\n", anahtar, deger)
	}

}
