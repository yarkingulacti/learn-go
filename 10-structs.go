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

Bu dosyada hem benzinli hem elektrikli motor örnekleri ve
interface kullanımı gösterilmektedir. Örneklerde sahip bilgisi
ile birlikte motorun kalan mil hesaplaması ve interface uyumu
gösterilir. 🚗🔋
*/
func yapilar() {
	var benimBenzinMotorum benzinMotoru = benzinMotoru{yakitVerimi: 30, galon: 10, sahipBilgisi: sahip{isim: "John"}, sayisal: 10}
	benimBenzinMotorum.yakitVerimi = 20
	var benimBenzinMotorum2 benzinMotoru = benzinMotoru{30, 10, sahip{"John"}, 10} // bildirime göre atama

	fmt.Println("🏷️ Benzinli motor örneği:", benimBenzinMotorum2)
	fmt.Printf("📣 %v'nin benzinli motoru tahmini %v mil gidebilir\n", benimBenzinMotorum2.sahipBilgisi.isim, benimBenzinMotorum2.yakitVerimi*benimBenzinMotorum2.galon)

	// anonim struct tanımı ve örneklendirme
	var benimElektrikMotorum = struct {
		batteryCapacity uint16
		sahipBilgisi    sahip
	}{400, sahip{"Alice"}}

	fmt.Println("🔋 Anonim elektrik motoru örneği:", benimElektrikMotorum)

	var benimBenzinMotorum3 benzinMotoru = benzinMotoru{yakitVerimi: 25, galon: 8, sahipBilgisi: sahip{isim: "Bob"}, sayisal: 5}
	fmt.Printf("📣 %v'nin benzinli motoru tahmini %v mil gidebilir (metod kullanılarak)\n", benimBenzinMotorum3.sahipBilgisi.isim, benimBenzinMotorum3.kalanMil())

	var mesafe uint8 = 200

	var benimBenzinMotorum4 benzinMotoru = benzinMotoru{yakitVerimi: 30, galon: 10, sahipBilgisi: sahip{isim: "John"}, sayisal: 10}
	var benimElektrikMotorum2 elektrikMotoru = elektrikMotoru{mpkwh: 4, kwh: 50}

	if yetisebilirMi(benimBenzinMotorum4, mesafe) {
		fmt.Printf("✅ %v'nin benzinli motoru %v mili gidebilir\n", benimBenzinMotorum4.sahipBilgisi.isim, mesafe)
	} else {
		fmt.Printf("❌ %v'nin benzinli motoru %v mili gidemez\n", benimBenzinMotorum4.sahipBilgisi.isim, mesafe)
	}

	if yetisebilirMi(benimElektrikMotorum2, mesafe) {
		fmt.Printf("✅ Elektrikli motor %v mili gidebilir\n", mesafe)
	} else {
		fmt.Printf("❌ Elektrikli motor %v mili gidemez\n", mesafe)
	}
}
