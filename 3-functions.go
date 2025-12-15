package main

import "fmt"

func fonksiyonlar() {
	// yaygın fonksiyon hata yakalama deseni
	var bolumSonucu, kalan, hata = tamsayiBolme(10, 0)

	if hata != nil {
		fmt.Printf("Error occurred during division %v", hata.Error())
	} else if kalan == 0 {
		fmt.Printf("The result of the integer division is %v", bolumSonucu)
	} else {
		fmt.Printf("The result of the integer division is %v with a remainder of %v", bolumSonucu, kalan)
	}

	switch {
	case hata != nil:
		fmt.Printf("Error occurred during division %v", hata.Error())
	case kalan == 0:
		fmt.Printf("The result of the integer division is %v", bolumSonucu)
	default:
		fmt.Printf("The result of the integer division is %v with a remainder of %v", bolumSonucu, kalan)
	}

	switch kalan {
	case 0:
		fmt.Printf("The result of the integer division is %v", bolumSonucu)
	case 1:
		fmt.Printf("The result of the integer division is %v with a remainder of 1", bolumSonucu)
	default:
		fmt.Printf("The result of the integer division is %v with a remainder of %v", bolumSonucu, kalan)
	}
}

func tamsayiBolme(bolunen, bolen int) (int, int, error) {
	if bolen == 0 {
		return 0, 0, fmt.Errorf("divisor cannot be zero")
	}
	var sonuc int = bolunen / bolen
	var kalan int = bolunen % bolen
	return sonuc, kalan, nil
}
