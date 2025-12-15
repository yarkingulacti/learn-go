package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Go'da Goroutineler:
  - Hafif iş parçacıklarıdır (lightweight threads)
  - `go` anahtar kelimesi ile başlatılır
  - Eşzamanlı (concurrent) programlama için kullanılır
  - Kanal (channel) ile iletişim kurabilirler
*/
func goRoutineler() {
	// t0 := time.Now()
	// for i := 0; i < len(dbVerisi); i++ {
	// 	wg.Add(1)
	// 	go dbCagrisi(i)
	// }

	// wg.Wait() // Tüm goroutinelerin bitmesini bekler
	// fmt.Printf("⏱️ Toplam süre (senkron): %v\n", time.Since(t0))
	// fmt.Printf("📊 DB Sonuçları: %v\n", dbSonuc)

	t0 := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go sahteDbCagrisi()
	}

	wg.Wait() // Tüm goroutinelerin bitmesini bekler
	fmt.Printf("⏱️ Toplam süre (1 milyon sahte çağrı): %v\n", time.Since(t0))
}

var rwM = sync.RWMutex{} // Okuma/Yazma kilidi
// var m = sync.Mutex{}      // Kritik bölge için mutex kilidi
var wg = sync.WaitGroup{} // WaitGroup eşzamanlama için sayaç görevi görür
var dbVerisi = []string{"id1", "id2", "id3", "id4", "id5"}
var dbSonuc = make([]string, len(dbVerisi))

func dbCagrisi(i int) {
	// Rastgele uyku süresi ile simüle edilmiş bir veri tabanı çağrısı
	var gecikme float32 = rand.Float32() * 1000
	time.Sleep(time.Duration(gecikme) * time.Millisecond)
	fmt.Printf("✅ DB çağrısı tamamlandı: %v\n", dbVerisi[i])
	// m.Lock() // Kritik bölgeyi kilitler, eğer başka bir goroutine erişmeye çalışırsa bekler
	kaydet(dbVerisi[i])
	logla()
	wg.Done() // İş tamamlandığında sayaçtan düşürür
}

func kaydet(deger string) {
	rwM.Lock()
	dbSonuc = append(dbSonuc, deger)
	rwM.Unlock()
}

func logla() {
	rwM.RLock()
	fmt.Println("📄 DB Sonuçları (log):", dbSonuc)
	rwM.RUnlock()
}

func sahteDbCagrisi() {
	// Basit bir sahte DB çağrısı örneği
	var gecikme float32 = 2000
	time.Sleep(time.Duration(gecikme) * time.Millisecond)
	wg.Done()
}
