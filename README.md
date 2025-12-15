# Go Öğreniyorum

Bu dosya Go (Golang) dilinin temel kavramlarını, kullanımını ve pratik örneklerini Türkçe olarak detaylı biçimde açıklar. Hedefi: yeni başlayanların ve temel bilgilerini tazelemek isteyenlerin ihtiyaç duyacağı açıklama ve örnekleri tek bir yerde toplamaktır.

## İçindekiler

- Giriş ve özellikler
- Kurulum ve araçlar
- Modüller ve paketler
- Temel tipler
- Değişkenler ve sabitler
- Kontrol akışları
- Fonksiyonlar
- Diziler, dilimler (slices) ve haritalar (maps)
- Struct'lar, metodlar ve arayüzler (interfaces)
- İşaretçiler (pointers)
- Hata yönetimi
- Eşzamanlılık (goroutines, channel, select)
- Test etme ve araçlar
- En iyi uygulamalar ve kaynaklar

---

## 1. Giriş ve Go'nun Temel Özellikleri

Go, Google tarafından geliştirilen, sistem programlama ve ağ servisleri için tasarlanmış, statik tipli ve derlenen bir dildir. Öne çıkan özellikleri:

- **Statik ve güçlü tip sistemi**: Derleme zamanında tip denetimi.
- **Hızlı derleme**: Büyük projelerde bile hızlı derleme süreleri.

# Learn Go 🚀

Bu küçük repo, Go dilinin temel kavramlarını basit ve eğlenceli örneklerle gösterir. Her konu için ayrı bir dosya var — amacımız hızlı öğrenme ve deneme yapmak.

Kısaca içerik:

- `main.go` : Program giriş noktası; örnek fonksiyonları çalıştırır.
- `1-variables.go` : Değişkenler ve temel tipler örnekleri.
- `2-conditions.go` : If / switch örnekleri.
- `3-functions.go` : Fonksiyonlar ve hata döndürme örnekleri.
- `4-arrays.go` : Diziler ve adresleri.
- `5-for-loop.go` : For döngülerine dair örnekler.
- `6-maps.go` : Haritalar (maps) ve kullanım örnekleri.
- `7-slices.go` : Dilimler (slices), `append`, `make`.
- `8-strings.go` : String, rune, `strings.Builder` örnekleri.
- `9-runes.go` : Rune'lar ve Unicode örnekleri.
- `10-structs.go` : Struct, metod ve interface örnekleri.
- `11-pointers.go` : İşaretçiler (pointers) ve bellek davranışı.

Nasıl çalıştırılır:

```powershell
cd C:\Users\yarkin\workspace\learn-go
go run .
```

Notlar ve ipuçları:

- Bu repo eğitim amaçlıdır; çıktılar Türkçe ve emojili olacak şekilde düzenlenmiştir.
- Kodları değiştirip denemekten çekinmeyin — her dosya kendi içinde bağımsız örnekler içerir.
- Daha fazlasını öğrenmek isterseniz `go` dokümantasyonuna bakabilirsiniz: https://go.dev

Katkı ve geliştirme:

- İstersen `fmt` çıktılarının dilini değiştirip kendi notlarını ekleyebilirsin.
- PR açmak istersen, branch oluşturup değişiklikleri gönder; ben yardımcı olurum.

Keyifli öğrenmeler! 🎉

- Tour of Go: https://tour.golang.org/

---

Bu rehberi geliştirmek veya belirli bir konu için örnekler eklememi isterseniz belirtin; örneğin `context` kullanımı, paket düzeni, ileri seviye concurrency desenleri veya test ve benchmark örnekleri ekleyebilirim.
