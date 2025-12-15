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
	// UTF-8 bayt indeksleme
	var benimMetin = "résumé"
	var indekslenen = benimMetin[0]
	fmt.Println(benimMetin)
	fmt.Println(indekslenen)
	fmt.Printf("%v, %T", indekslenen, indekslenen)
	for i, v := range benimMetin {
		fmt.Printf("Index: %v, Value: %v", i, v)
	}

	fmt.Println("\nThe length of 'myString' is %v", len(benimMetin))

	var metinDilimi []string = []string{"r", "é", "s", "u", "m", "é"}
	var birlesikMetin string

	for i := range metinDilimi {
		// her seferinde yeni string oluşturulur
		birlesikMetin += metinDilimi[i]
	}

	fmt.Println(birlesikMetin)

	var metinYapici strings.Builder
	var metinDilimi2 []string = []string{"r", "é", "s", "u", "m", "é"}

	for i := range metinDilimi2 {
		metinYapici.WriteString(metinDilimi2[i])
	}

	fmt.Println(metinYapici.String())
}
