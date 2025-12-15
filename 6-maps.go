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
	// Harita oluşturma (örnekler ve açıklamalar)
	benimHarita := make(map[string]uint8)
	fmt.Println("🗺️ Boş harita örneği:", benimHarita)

	// Atama sırasında anahtar-değer çifti ekleme
	benimHarita2 := map[string]uint8{"Adam": 30, "Eve": 28}
	fmt.Println("🎯 'Adam' anahtarına karşılık gelen yaş:", benimHarita2["Adam"])
	fmt.Println("ℹ️ 'Sarah' anahtarı yoksa varsayılan değer:", benimHarita2["Sarah"]) // varsayılan değer 0

	yas, varMi := benimHarita2["Sarah"]

	if varMi {
		fmt.Println("✅ Sarah bulundu, yaşı:", yas)
	} else {
		fmt.Println("❌ Anahtar bulunamadı: Sarah")
	}

	// bir anahtar-değer çiftini silme
	delete(benimHarita2, "Adam")

	fmt.Println("🗺️ Harita (Adam silindikten sonra):", benimHarita2)

	for anahtar, deger := range benimHarita2 {
		fmt.Printf("👤 İsim: %v, 🔢 Yaş: %v\n", anahtar, deger)
	}

}
