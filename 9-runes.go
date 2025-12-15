package main

import "fmt"

/*
Runeler (Runes) Go'da:
	- int32 takma adıdır, bir Unicode kod noktasını temsil eder
	- Bireysel runelere indeksleme ile erişilebilir
	- ASCII dışındaki karakterleri işlemek için yararlıdır
*/
func runeOrnekleri() {
	// UTF-8 rune indeksleme ve rune kullanım örnekleri
	var benimRuneDizisi = []rune{'r', 'é', 's', 'u', 'm', 'é'}
	var indekslenen = benimRuneDizisi[0]
	fmt.Println("🔣 Rune dizisi örneği:", benimRuneDizisi)
	fmt.Println("🔢 İlk rune (kod noktası):", indekslenen)
	fmt.Printf("🔎 Tip: %v, %T\n", indekslenen, indekslenen)
	fmt.Println("🔁 Rune bazlı yineleme:")
	for i, v := range benimRuneDizisi {
		fmt.Printf("  • İndeks: %v, Rune: %v\n", i, v)
	}

	var benimRune = 'a'
	fmt.Printf("📌 Tek rune örneği: %v\n", benimRune)
}
