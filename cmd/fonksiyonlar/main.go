package fonksiyonlar

import "fmt"

func main() {
	FonksiyonlarIcerik()
}

func FonksiyonlarIcerik() {
	// defer ifadesi, fonksiyonun sonunda veya return ifadesinden hemen önce çalıştırılacak kodu belirtir
	defer fmt.Println("Fonksiyonlar paketinden çıkılıyor...")

	// Yaygın fonksiyon hata yakalama deseni (açıklamalı)
	// fonksiyonlara parametre olarak fonksiyonlar da geçilebilir
	bolen, bolunen := 2, 10
	bolumSonucu, kalan, hata := tamsayiBolme(bolunen, bolen, kalaniAl, sonucuAl)

	if hata != nil {
		fmt.Printf("❗ Hata: Bölme sırasında hata oluştu: %v\n", hata.Error())
	} else if kalan == 0 {
		fmt.Printf("✅ Tam bölme: Bölüm sonucu = %v\n", bolumSonucu)
	} else {
		fmt.Printf("🔢 Bölme sonucu = %v, kalan = %v\n", bolumSonucu, kalan)
	}

	// aynı durumu switch ile gösterme örneği
	switch {
	case hata != nil:
		fmt.Printf("❗ Hata (switch): %v\n", hata.Error())
	case kalan == 0:
		fmt.Printf("✅ Tam bölme (switch): %v\n", bolumSonucu)
	default:
		fmt.Printf("🔢 Bölme (switch) sonucu = %v, kalan = %v\n", bolumSonucu, kalan)
	}

	// kalan değerine göre farklı mesaj
	switch kalan {
	case 0:
		fmt.Printf("🎯 Kalan yok: sonuç = %v\n", bolumSonucu)
	case 1:
		fmt.Printf("🔔 Kalan 1: sonuç = %v, kalan = 1\n", bolumSonucu)
	default:
		fmt.Printf("🔎 Genel durum: sonuç = %.2f, kalan = %d\n", bolumSonucu, kalan)
	}

	// _ kullanarak istenmeyen dönüş değerlerini yoksayma
	var selamMetni, _ = baslikVeAltCizgi()

	fmt.Println(selamMetni)
	fmt.Println(koordinatlariAl())
	// fonksiyona anonim fonksiyon parametresi geçme örneği
	sonuc := isaretle(func(metin string) string {
		return "~~ " + metin + " ~~"
	}, "Merhaba Dünya")

	fmt.Println(sonuc)

	// kapanış fonksiyonu örneği
	topla := toplayici()
	toplam1 := topla(5)
	toplam2 := topla(10)
	toplam3 := topla(20)

	fmt.Printf("Kapanış fonksiyonu sonuçları: %d, %d, %d\n", toplam1, toplam2, toplam3)
}

func tamsayiBolme(bolunen, bolen int, kalaniAl func(int, int) int, sonucuAl func(int, int) float64) (float64, int, error) {
	// bolunen ve bolen parametrelerinin her ikisi de int türündedir.

	if bolen == 0 {
		return 0, 0, fmt.Errorf("bölen sıfır olamaz")
	}

	sonuc := sonucuAl(bolunen, bolen)
	kalan := kalaniAl(bolunen, bolen)

	return sonuc, kalan, nil
}

func sonucuAl(bolunen, bolen int) float64 {
	return float64(bolunen) / float64(bolen)
}

func kalaniAl(bolunen, bolen int) int {
	return bolunen % bolen
}

func baslikVeAltCizgi() (string, string) {
	return "selam!", "merhaba!"
}

/*
Eğer bir fonksiyonun dönüş değerleri isimlendirilmişse, fonksiyonun içinde return ifadesi tek başına kullanıldığında, bu isimlendirilmiş dönüş değerlerinin varsayılan değerleri döndürülür. Dönüş değerlerini isimlendirmek, kodun okunabilirliğini artırabilir ve özellikle birden fazla dönüş değeri olduğunda hangi değerin ne anlama geldiğini daha açık hale getirebilir.
*/
func koordinatlariAl() (x, y, z int) {
	return
}

func isaretle(isaretci func(string) string, metin string) string {
	return isaretci(metin)
}

// Kapanış fonksiyonu örneği
// toplayici fonksiyonu, her çağrıldığında verilen sayıyı birikimli olarak toplar ve toplamı döndürür. bunun sebebi iç içe tanımlanan anonim fonksiyonun, dış fonksiyonun kapsamındaki 'toplam' değişkenine erişebilmesidir.
func toplayici() func(int) int {
	var toplam int

	return func(sayi int) int {
		toplam += sayi

		return toplam
	}
}
