package main

import (
	"fmt"
	"strings"
)

/*
Metinler (Strings) Go'da:
  - Baytların değiştirilemez sıraya dizilişi
  - Varsayılan olarak UTF-8 kodlamalıdır
  - Bireysel baytlara indeksleme ile erişilebilir
  - Uzunluk `len()` ile elde edilebilir
  - Birleştirme `+` operatörü veya verimli kullanım için `strings.Builder` ile yapılır
*/
func metinler() {
	// Bu fonksiyon Go'da string'lerin nasıl davrandığını gösterir.
	// UTF-8, byte vs rune farkları ve performans için strings.Builder kullanımı açıklanır.
	var benimMetin = "résumé"
	var indekslenen = benimMetin[0]
	fmt.Println("💬 Örnek metin:", benimMetin)
	fmt.Println("🔢 İlk bayt (byte) değeri:", indekslenen)
	fmt.Printf("🔎 Tip kontrolü: %v, %T\n", indekslenen, indekslenen)
	fmt.Println("🔁 Rune bazlı yineleme (index, rune):")
	for i, v := range benimMetin {
		fmt.Printf("  • İndeks: %v, Rune değeri: %v\n", i, v)
	}

	fmt.Printf("📏 Metin uzunluğu (bayt): %v\n", len(benimMetin))

	var metinDilimi []string = []string{"r", "é", "s", "u", "m", "é"}
	var birlesikMetin string

	for i := range metinDilimi {
		// her seferinde yeni string oluşturulur (performans maliyeti)
		birlesikMetin += metinDilimi[i]
	}

	fmt.Println("✂️ Elle birleştirme sonucu:", birlesikMetin)

	var metinYapici strings.Builder
	var metinDilimi2 []string = []string{"r", "é", "s", "u", "m", "é"}

	for i := range metinDilimi2 {
		metinYapici.WriteString(metinDilimi2[i])
	}

	fmt.Println("🚀 strings.Builder ile birleştirme:", metinYapici.String())
}
