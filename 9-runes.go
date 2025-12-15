package main

import "fmt"

/*
Runeler (Runes) Go'da:
	- int32 takma adıdır, bir Unicode kod noktasını temsil eder
	- Bireysel runelere indeksleme ile erişilebilir
	- ASCII dışındaki karakterleri işlemek için yararlıdır
*/
func runeOrnekleri() {
	// UTF-8 rune indeksleme
	var benimRuneDizisi = []rune{'r', 'é', 's', 'u', 'm', 'é'}
	var indekslenen = benimRuneDizisi[0]
	fmt.Println(benimRuneDizisi)
	fmt.Println(indekslenen)
	fmt.Printf("%v, %T", indekslenen, indekslenen)
	for i, v := range benimRuneDizisi {
		fmt.Printf("Index: %v, Value: %v", i, v)
	}

	var benimRune = 'a'
	fmt.Printf("\nmyRune = %v", benimRune)
}
