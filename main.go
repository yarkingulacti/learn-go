package main // başlangıç paket adı

import (
	"fmt"
	"learn_go/cmd/degiskenler"
	"learn_go/cmd/fonksiyonlar"
	"learn_go/cmd/kosullar"
	"learn_go/cmd/yapilar"
)

func main() {
	fmt.Println("🔢 === Değişkenler (Variables) ===")
	degiskenler.DegiskenlerIcerik()
	fmt.Println("⚖️ === Koşullar (Conditions) ===")
	kosullar.KosullarIcerik()
	fmt.Println("🧩 === Fonksiyonlar (Functions) ===")
	fonksiyonlar.FonksiyonlarIcerik()
	// fmt.Println("🧵 === Diziler (Arrays) ===")
	// diziler()
	// fmt.Println("🔁 === Döngüler (Loops) ===")
	// dongu()
	// fmt.Println("🗺️ === Haritalar (Maps) ===")
	// haritalar()
	// fmt.Println("📦 === Dilimler (Slices) ===")
	// dilimler()
	// fmt.Println("🔣 === Runeler (Runes) ===")
	// runeOrnekleri()
	fmt.Println("🏗️ === Yapılar (Structs) ===")
	yapilar.YapilarIcerik()
	// fmt.Println("📍 === İşaretçiler (Pointers) ===")
	// isaretciler()
	// fmt.Println("🚀 === Go Rutinler (Goroutines) ===")
	// goRoutineler()
	// fmt.Println("📡 === Kanallar (Channels) ===")
	// kanallar()
	// fmt.Println("🔢 === Generics (Tür Parametreleri) ===")
	// turParametreleri()
}
