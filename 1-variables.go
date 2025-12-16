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
	// Bu fonksiyon Go'nun temel değişken tiplerini gösterir.
	// Her örnek için kısa bir açıklama ve emoji ile kullanıcıya bilgi veriyoruz.
	var tamsayi int16 = 32767
	fmt.Println("🔢 Tam sayı örneği (int16):", tamsayi)

	var ondalikAzHassas float32 = 12345678.9 // hassasiyet kaybı
	var ondalikHassas float64 = 12345678.9   // daha hassas

	fmt.Println("🔬 float32 (az hassas):", ondalikAzHassas)
	fmt.Println("🔬 float64 (daha hassas):", ondalikHassas)

	var ondalik32 float32 = 10.1
	var tamsayi32 int32 = 2
	var sonuc float32 = ondalik32 + float32(tamsayi32) // doğru dönüşüm

	fmt.Println("➕ Tip dönüşümü örneği, sonuç:", sonuc)

	var benimMetin string = "Hello, \nGo!"
	fmt.Println("💬 String örneği (escape karakterleri gösterilir):")
	fmt.Println(benimMetin)

	fmt.Println("🧾 'ğ' karakterinin byte uzunluğu (len):", len("ğ"))
	fmt.Println("🔠 'ğ' karakterinin rune sayısı (utf8):", utf8.RuneCountInString("ğ"))

	var benimRune rune = 'a' // ASCII'de 97
	fmt.Println("🔣 Rune örneği (kod noktası):", benimRune)

	var dogruMu bool = true
	fmt.Println("✅ Boolean örneği (true):", dogruMu)

	var digerBool bool // varsayılan değer false
	fmt.Println("❌ Boolean varsayılan değeri (false):", digerBool)

	var benimMetin2 string
	fmt.Println("🔤 String varsayılan değeri (boş):", benimMetin2)

	var benimRune2 rune
	fmt.Println("0 değeri olan rune (default):", benimRune2)

	var metinim = "text" // türü string olarak çıkarılır
	fmt.Println("✳️ Tür çıkarmaya örnek (string):", metinim)

	kisaDegisken := "kısa değişken beyanı" // tür çıkarımı(type inference)
	fmt.Println("📝 Kısa değişken beyanı örneği:", kisaDegisken)

	var sayi1, sayi2, sayi3 int = 1, 2, 3
	fmt.Println("1,2,3 örneği (çoklu atama):", sayi1, sayi2, sayi3)

	const benimSabit = "Bu bir sabittir; değeri değiştirilemez ve tanımlanırken başlatılmalıdır."
	fmt.Println("📌 Sabit örneği:", benimSabit)

	fmt.Printf("Değişkenleri gösterme;\n1. Normal interpolasyon örneği: %v\n", benimMetin)
	fmt.Printf("2. Metin interpolasyonu ile: %s\n", benimMetin)
	fmt.Printf("3. Tamsayı interpolasyonu ile: %d\n", tamsayi)
	fmt.Printf("4. Ondalık interpolasyonu ile: %f\n", ondalikHassas)
	fmt.Printf("5. Ondalık(formatlı) interpolasyonu ile: %.2f\n", ondalikHassas)
	fmt.Printf("6. Bool interpolasyonu ile: %t\n", dogruMu)

	const isim = "Yarkın"
	const soyisim = "Gülaçtı"
	tamIsim := fmt.Sprintf("%s %s", isim, soyisim)

	fmt.Println("👤 Tam isim oluşturma (Sprintf ile):", tamIsim)
}
