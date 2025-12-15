package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Go'da Kanallar (Channels):
  - Go rutinleri arasında veri iletimi için kullanılır
  - `chan` anahtar kelimesi ile tanımlanır
  - Gönderme ve alma işlemleri için `<-` operatörü kullanılır
  - Eşzamanlı programlamada iletişim için temel yapı taşıdır
*/
func kanallar() {
	var kanal = make(chan int)
	// Yanlış kullanım örneği (bloklanma nedeniyle program kilitlenir)
	// kanal <- 10
	// deger := <-kanal
	// fmt.Println("📡 Kanal üzerinden alınan değer:", deger)
	go islem(kanal)

	fmt.Println(<-kanal)

	go donguselIslem(kanal)

	for i := 0; i < 5; i++ {
		fmt.Printf("📡 Kanal üzerinden alınan değer: %v\n", <-kanal)
	}

	for deger := range kanal {
		fmt.Printf("📡 Kanal üzerinden alınan değer: %v\n", deger)
	}

	var bufferedKanal = make(chan int, 5)

	go donguselIslem(bufferedKanal)

	for deger := range bufferedKanal {
		fmt.Printf("📡 Buffered Kanal üzerinden alınan değer: %v\n", deger)
		time.Sleep(1 * time.Second)
	}

	var tavukKanal = make(chan string)
	var etKanal = make(chan string)
	var saticilar = []string{"Tavukçu Ahmet", "Köfteci Mehmet", "Kanatçı Ayşe"}

	for i := range saticilar {
		go tavukFiyatlariniKontrolEt(saticilar[i], tavukKanal)
		go etFiyatlariniKontrolEt(saticilar[i], etKanal)
	}

	mesajGonder(tavukKanal, etKanal)
}

var MAX_TAVUK_FIYATI float32 = 5
var MAX_ET_FIYATI float32 = 10

func islem(k chan int) {
	k <- 123
}

func donguselIslem(k chan int) {
	defer close(k) // Kanalı kapatmayı fonksiyon sonunda garanti eder

	for i := 0; i < 5; i++ {
		k <- i
	}

	fmt.Println("✅ Gönderim tamamlandı, kanal kapatılıyor")
}

func tavukFiyatlariniKontrolEt(satici string, tavukKanal chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tavukFiyati = rand.Float32() * 20
		if tavukFiyati <= MAX_TAVUK_FIYATI {
			tavukKanal <- satici
			break
		}
	}
}

func etFiyatlariniKontrolEt(satici string, etKanal chan string) {
	for {
		time.Sleep(time.Second * 1)
		var etFiyati = rand.Float32() * 20
		if etFiyati <= MAX_ET_FIYATI {
			etKanal <- satici
			break
		}
	}
}

func mesajGonder(tavukKanal chan string, etKanal chan string) {
	select {
	case tavukSatici := <-tavukKanal:
		fmt.Printf("🍗 Tavuk alımı için en uygun satıcı: %v\n", tavukSatici)
	case etSatici := <-etKanal:
		fmt.Printf("🍖 Et alımı için en uygun satıcı: %v\n", etSatici)
	}
}
