package main

import "fmt"

func fonksiyonlar() {
	// Yaygın fonksiyon hata yakalama deseni (açıklamalı)
	var bolumSonucu, kalan, hata = tamsayiBolme(10, 0)

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
		fmt.Printf("🔎 Genel durum: sonuç = %v, kalan = %v\n", bolumSonucu, kalan)
	}
}

func tamsayiBolme(bolunen, bolen int) (int, int, error) {
	if bolen == 0 {
		return 0, 0, fmt.Errorf("bölen sıfır olamaz")
	}
	var sonuc int = bolunen / bolen
	var kalan int = bolunen % bolen
	return sonuc, kalan, nil
}
