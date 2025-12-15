package main

import "fmt"

type benzinMotoru struct {
	yakitVerimi  uint8
	galon        uint8
	sahipBilgisi sahip // benzinMotoru.sahipBilgisi.isim
	sayisal      int   // default değer 0
}

type elektrikMotoru struct {
	mpkwh uint8
	kwh   uint8
}

type motor interface {
	kalanMil() uint8
}

type sahip struct {
	isim string
}

// benzinMotoru tipi için bir metod tanımlıyoruz
func (b benzinMotoru) kalanMil() uint8 {
	return b.yakitVerimi * b.galon
}

func (e elektrikMotoru) kalanMil() uint8 {
	return e.mpkwh * e.kwh
}

func yetisebilirMi(m motor, mesafe uint8) bool {
	return m.kalanMil() >= mesafe
}

/*
Yapılar (Structs) Go'da:
	- İlgili alanları gruplayan özel veri tipleri
	- 'type' ve 'struct' anahtar kelimeleri ile tanımlanır
	- Alanlara nokta gösterimi ile erişilir ve değiştirilebilir
	- Metod desteği ile davranış eklenebilir
*/
func yapilar() {
	var benimBenzinMotorum benzinMotoru = benzinMotoru{yakitVerimi: 30, galon: 10, sahipBilgisi: sahip{isim: "John"}, sayisal: 10}
	benimBenzinMotorum.yakitVerimi = 20
	var benimBenzinMotorum2 benzinMotoru = benzinMotoru{30, 10, sahip{"John"}, 10} // bildirime göre atama

	fmt.Println(benimBenzinMotorum2)
	fmt.Printf("%v's gas engine can go %v miles\n", benimBenzinMotorum2.sahipBilgisi.isim, benimBenzinMotorum2.yakitVerimi*benimBenzinMotorum2.galon)

	// anonim struct tanımı ve örneklendirme
	var benimElektrikMotorum = struct {
		batteryCapacity uint16
		sahipBilgisi    sahip
	}{400, sahip{"Alice"}}

	fmt.Println(benimElektrikMotorum)

	var benimBenzinMotorum3 benzinMotoru = benzinMotoru{yakitVerimi: 25, galon: 8, sahipBilgisi: sahip{isim: "Bob"}, sayisal: 5}
	fmt.Printf("%v's gas engine can go %v miles\n", benimBenzinMotorum3.sahipBilgisi.isim, benimBenzinMotorum3.kalanMil())

	var mesafe uint8 = 200

	var benimBenzinMotorum4 benzinMotoru = benzinMotoru{yakitVerimi: 30, galon: 10, sahipBilgisi: sahip{isim: "John"}, sayisal: 10}
	var benimElektrikMotorum2 elektrikMotoru = elektrikMotoru{mpkwh: 4, kwh: 50}

	if yetisebilirMi(benimBenzinMotorum4, mesafe) {
		fmt.Printf("%v's gas engine can make it %v miles\n", benimBenzinMotorum4.sahipBilgisi.isim, mesafe)
	} else {
		fmt.Printf("%v's gas engine cannot make it %v miles\n", benimBenzinMotorum4.sahipBilgisi.isim, mesafe)
	}

	if yetisebilirMi(benimElektrikMotorum2, mesafe) {
		fmt.Printf("The electric engine can make it %v miles\n", mesafe)
	} else {
		fmt.Printf("The electric engine cannot make it %v miles\n", mesafe)
	}
}
