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

	fmt.Printf("The value p points to is: %v\n", *isaretciP) // p'nin işaret ettiği değeri çözme, int32'nin sıfır değeri 0
	fmt.Printf("The value of i is: %v\n", sayi)

	*isaretciP = 42 // p'nin işaret ettiği adresteki değeri ayarlama

	fmt.Printf("After setting, the value p points to is: %v\n", *isaretciP)

	// var p2 *int32 // nil pointer
	// *p2 = 21      // Bu runtime panic'e neden olur: nil işaretçi çözümleniyor

	fmt.Printf("Before setting, the value of i is: %v\n", sayi)
	var isaretciP3 *int32 = &sayi // p3 sayının adresine işaret eder
	*isaretciP3 = 21              // p3'ün işaret ettiği adresteki değeri ayarlama, yani sayi
	fmt.Printf("After setting through pointer, the value of i is: %v\n", sayi)

	var k int32 = 55
	var j int32

	fmt.Printf("Before Value of k: %v\n", k) // k 55
	fmt.Printf("Before Value of j: %v\n", j) // j 0

	j = 99 // j'yi değiştirmek k'yi etkilemez

	fmt.Printf("After Value of k: %v\n", k) // k aynı kalır
	fmt.Printf("After Value of j: %v\n", j) // j şimdi 99

	var dilim = []int32{1, 2, 3}            // dilimler referans tiplerdir
	fmt.Printf("Before Slice: %v\n", dilim) // [1 2 3]
	var dilimKopya = dilim                  // dilim ve dilimKopya aynı alt diziyi işaret eder
	dilimKopya[0] = 99                      // dilimKopya'yı değiştirmek dilimi etkiler

	fmt.Printf("After Slice: %v\n", dilim) // [99 2 3]

	var dizi = [3]int32{1, 2, 3}           // diziler değer tipidir
	fmt.Printf("Before Array: %v\n", dizi) // [1 2 3]
	var diziKopya = dizi                   // diziKopya dizi'nin kopyasıdır
	diziKopya[0] = 99                      // diziKopya'yı değiştirmek diziyi etkilemez

	fmt.Printf("After Array: %v\n", dizi) // [1 2 3]

	var yaslar = map[string]int32{
		"Alice": 30,
		"Bob":   25,
	} // haritalar referans tiplerdir
	fmt.Printf("Before Map: %v\n", yaslar)
	var yaslarKopya = yaslar  // yaslar ve yaslarKopya aynı veriyi gösterir
	yaslarKopya["Alice"] = 31 // yaslarKopya'yı değiştirmek yaslar'ı etkiler
	fmt.Printf("After Map: %v\n", yaslar)

	var nesne1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Before square: %v\n", nesne1)
	fmt.Printf("The memory location of thing1 array is: %p\n", &nesne1)
	var sonuc = kare(nesne1)
	fmt.Printf("After square: %v\n", sonuc)
	fmt.Printf("The value of thing1 is: %v", nesne1)

	var nesne2 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nBefore squarePointer: %v\n", nesne2)
	fmt.Printf("The memory location of thing2 array is: %p\n", &nesne2)
	var sonuc2 = kareIsaretci(&nesne2)
	fmt.Printf("After squarePointer: %v\n", sonuc2)
	fmt.Printf("The value of thing2 is: %v", nesne2)
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
