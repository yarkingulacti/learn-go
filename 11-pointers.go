package main

import "fmt"

/*
Pointers (İşaretçiler) Go'da:
	- Başka bir değişkenin bellek adresini tutan değişkenlerdir
	- '*' operatörü ile tanımlanır
	- '&' operatörü bir değişkenin adresini almak için kullanılır
	- '*' operatörü bir işaretçinin gösterdiği adresin değerine erişmek için kullanılır
	- Bellek verimliliği ve fonksiyonlarda değişkenleri değiştirmek için faydalıdır
*/
func isaretciler() {
	// var p *int32 // başlangıçta nil çünkü herhangi bir adrese işaret etmiyor
	var isaretciP *int32 = new(int32) // int32 için bellek ayırır ve adresini döndürür
	var sayi int32 = 31

	fmt.Printf("📍 isaretciP'nin işaret ettiği değer (başlangıç, sıfır): %v\n", *isaretciP) // p'nin işaret ettiği değeri çözme, int32'nin sıfır değeri 0
	fmt.Printf("🔢 sayi değişkeninin değeri: %v\n", sayi)

	*isaretciP = 42 // p'nin işaret ettiği adresteki değeri ayarlama

	fmt.Printf("✳️ isaretciP ile ayarlanan yeni değer: %v\n", *isaretciP)

	// var p2 *int32 // nil pointer
	// *p2 = 21      // Bu runtime panic'e neden olur: nil işaretçi çözümleniyor

	fmt.Printf("🔁 İşaretçi ile atama öncesi sayi: %v\n", sayi)
	var isaretciP3 *int32 = &sayi // p3 sayının adresine işaret eder
	*isaretciP3 = 21              // p3'ün işaret ettiği adresteki değeri ayarlama, yani sayi
	fmt.Printf("🔁 İşaretçi ile atama sonrası sayi: %v\n", sayi)

	var k int32 = 55
	var j int32

	fmt.Printf("🔎 k değişkeninin başlangıç değeri: %v\n", k) // k 55
	fmt.Printf("🔎 j değişkeninin başlangıç değeri: %v\n", j) // j 0

	j = 99 // j'yi değiştirmek k'yi etkilemez

	fmt.Printf("🔁 k değişkeni (değişmedi): %v\n", k) // k aynı kalır
	fmt.Printf("🔁 j değişkeni (güncel): %v\n", j)    // j şimdi 99

	var dilim = []int32{1, 2, 3}              // dilimler referans tiplerdir
	fmt.Printf("📌 Dilim (önce): %v\n", dilim) // [1 2 3]
	var dilimKopya = dilim                    // dilim ve dilimKopya aynı alt diziyi işaret eder
	dilimKopya[0] = 99                        // dilimKopya'yı değiştirmek dilimi etkiler

	fmt.Printf("📌 Dilim (sonra, referans tipi nedeniyle değişti): %v\n", dilim) // [99 2 3]

	var dizi = [3]int32{1, 2, 3}            // diziler değer tipidir
	fmt.Printf("📌 Dizi (önce): %v\n", dizi) // [1 2 3]
	var diziKopya = dizi                    // diziKopya dizi'nin kopyasıdır
	diziKopya[0] = 99                       // diziKopya'yı değiştirmek diziyi etkilemez

	fmt.Printf("📌 Dizi (sonra, kopya değişti ama orijinal aynı kaldı): %v\n", dizi) // [1 2 3]

	var yaslar = map[string]int32{
		"Alice": 30,
		"Bob":   25,
	} // haritalar referans tiplerdir
	fmt.Printf("🗺️ Harita (önce): %v\n", yaslar)
	var yaslarKopya = yaslar  // yaslar ve yaslarKopya aynı veriyi gösterir
	yaslarKopya["Alice"] = 31 // yaslarKopya'yı değiştirmek yaslar'ı etkiler
	fmt.Printf("🗺️ Harita (sonra, referans tipi nedeniyle değişti): %v\n", yaslar)

	var nesne1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("🔢 Dizi örneği (önce): %v\n", nesne1)
	fmt.Printf("📍 nesne1 dizisinin bellek adresi: %p\n", &nesne1)
	var sonuc = kare(nesne1)
	fmt.Printf("🔁 kare() sonrası (kopya): %v\n", sonuc)
	fmt.Printf("🔎 nesne1 orijinali (değişmedi): %v\n", nesne1)

	var nesne2 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n🔢 Dizi örneği (pointer ile önce): %v\n", nesne2)
	fmt.Printf("📍 nesne2 dizisinin bellek adresi: %p\n", &nesne2)
	var sonuc2 = kareIsaretci(&nesne2)
	fmt.Printf("🔁 kareIsaretci() sonrası (yerinde değişiklik): %v\n", sonuc2)
	fmt.Printf("🔎 nesne2 orijinali (değişti): %v\n", nesne2)
}

// Bu şekilde fonksiyon içinde dizinin yeni belleğini ayırıp yeni bir dizi döndürüyoruz
// ve orijinal dizi değişmeden kalıyor.
// Bu, dizi büyükse daha fazla bellek kullanımıyla sonuçlanabilir.
// Bunun yerine dizinin işaretçisini geçirip yerinde değiştirebiliriz.
func kare(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 array in square func is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return thing2
}

// Bu şekilde dizinin işaretçisini geçiriyoruz, orijinal diziyi yerinde değiştirmemizi sağlar
// ve kopya için yeni bellek ayırmaz. Bu özellikle büyük diziler için daha verimlidir.
func kareIsaretci(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 array in squarePointer func is: %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}
