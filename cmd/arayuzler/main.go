package arayuzler

import "fmt"

func main() {
	ArayuzlerIcerik()
}

type oval interface {
	cap() float64
}

type sekil interface {
	alan() float64
	cevre() float64
	alanMaliyeti(float64) float64
	cevreMaliyeti(fiyat float64) float64
}

type dikdortgen struct {
	uzunluk, genislik float64
}

func (d dikdortgen) alan() float64 {
	return d.uzunluk * d.genislik
}

func (d dikdortgen) cevre() float64 {
	return 2 * (d.uzunluk + d.genislik)
}

func (d dikdortgen) alanMaliyeti(fiyat float64) float64 {
	return d.alan() * fiyat
}

func (d dikdortgen) cevreMaliyeti(fiyat float64) float64 {
	return d.cevre() * fiyat
}

type daire struct {
	yariCap float64
}

func (d daire) alan() float64 {
	return 3.14 * d.yariCap * d.yariCap
}

func (d daire) cevre() float64 {
	return float64(2) * float64(3.14) * d.yariCap
}

func (d daire) alanMaliyeti(fiyat float64) float64 {
	return d.alan() * fiyat
}

func (d daire) cevreMaliyeti(fiyat float64) float64 {
	return d.cevre() * fiyat
}

func (d daire) cap() float64 {
	return d.yariCap * 2
}

/*
Go'da Arayüzler (Interfaces):
  - Method setleri
  - Tür uyumluluğu sağlar
*/
func ArayuzlerIcerik() {
	var s sekil
	var o oval

	c := daire{yariCap: 10}
	s = c
	o = c
	fmt.Println("Dairenin Alanı:", s.alan())
	fmt.Println("Dairenin Çevresi:", s.cevre())
	fmt.Println("Dairenin Çapı:", o.cap())
	fmt.Println("Dairenin Alan Maliyeti:", s.alanMaliyeti(5))
	fmt.Println("Dairenin Çevre Maliyeti:", s.cevreMaliyeti(3))

	d := dikdortgen{uzunluk: 5, genislik: 3}
	s = d
	fmt.Println("Dikdörtgen Alan:", s.alan())
	fmt.Println("Dikdörtgen Çevre:", s.cevre())
	fmt.Println("Dikdörtgen Alan Maliyeti:", s.alanMaliyeti(4))
	fmt.Println("Dikdörtgen Çevre Maliyeti:", s.cevreMaliyeti(2))

	// Tip iddiası (type assertion)
	c, ok := s.(daire)

	if ok {
		fmt.Println("s değişkeni daire türündedir. Yarıçapı:", c.yariCap)
	} else {
		fmt.Println("s değişkeni daire türünde değildir.")
	}

	daireOrnegi := interface{}(daire{yariCap: 7})

	// Tip switch
	switch v := daireOrnegi.(type) {
	case sekil:
		fmt.Printf("dairenin çevresi: %f\n", v.cevre())
	default:
		fmt.Printf("sayi bilinmeyen bir %T türündedir.\n", v)
	}
}
