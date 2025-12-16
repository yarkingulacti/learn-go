package main

import "fmt"

func kosullar() {
	if 5 > 3 {
		fmt.Println("👍 Örnek karşılaştırma: 5, 3'ten büyüktür")
	} else {
		fmt.Println("👎 Örnek karşılaştırma: 5, 3'ten büyük değildir")
	}

	if rakam := 7; rakam%2 == 0 {
		fmt.Println("🔢 Rakam kontrolü: rakam çifttir (even)")
	} else {
		fmt.Println("🔢 Rakam kontrolü: rakam tektir (odd)")
	}

	var sayi int = 10

	if sayi%2 == 0 {
		fmt.Println("🔢 Sayı kontrolü: sayı çifttir (even)")
	} else {
		fmt.Println("🔢 Sayı kontrolü: sayı tektir (odd)")
	}

	switch sayi {
	case 1:
		fmt.Println("🔎 Durum: sayı 1'dir")
	case 5:
		fmt.Println("🔎 Durum: sayı 5'tir")
	case 10:
		fmt.Println("🔎 Durum: sayı 10'dur")
	default:
		fmt.Println("🔎 Durum: sayı listede değil")
	}

	switch {
	case sayi < 0:
		fmt.Println("⚠️ Sayı negatif")
	case sayi == 0:
		fmt.Println("⚠️ Sayı sıfır")
	case sayi > 0:
		fmt.Println("✅ Sayı pozitif")
	}
}
