# Go Öğreniyorum — Basit ve Açıklayıcı Rehber

Bu depo, Go (Golang) dilini öğrenenler için hazırlanmış küçük, pratik ve örnek ağırlıklı bir çalışma dizinidir. Amaç: karmaşık terimlere boğmadan, hızlıca deneme yapabileceğiniz, anlamayı kolaylaştıran açıklamalar sunmaktır.

Not: Yazılım jargonunu korudum (ör. `goroutine`, `channel`, `interface`, `iota`, `go mod`), ancak açıklamaları Türkçe ve sade tuttum.

## Hızlı Bakış — Dosyalar

- `main.go` — Programın giriş noktası; burada diğer örnek fonksiyonlar sırayla çalıştırılır.
- `1-variables.go` — Değişkenler, tipler, sabitler ve örnek kullanım.
- `2-conditions.go` — `if` ve `switch` örnekleri.
- `3-functions.go` — Fonksiyonlar, çoklu dönüş, hata döndürme örnekleri.
- `4-arrays.go` — Sabit uzunluklu diziler ve bellek adresleri.
- `5-for-loop.go` — `for` ile yapılan yinelemeler ve farklı for formları.
- `6-maps.go` — Map (harita) kullanımı, okuma, yazma ve silme.
- `7-slices.go` — Slice, `append`, `make`, length/capacity açıklamaları.
- `8-strings.go` — `string`, `rune`, byte vs rune farkı, `strings.Builder`.
- `9-runes.go` — Rune örnekleri ve Unicode işleme.
- `10-structs.go` — `struct`, metotlar ve `interface` kullanımı.
- `11-pointers.go` — Pointer (işaretçi) davranışı, değer vs referans farkı.

## Nasıl Çalıştırılır

1. Komut satırında proje dizinine gidin:

```powershell
cd C:\Users\yarkin\workspace\learn-go
```

2. Programı çalıştırın:

```powershell
go run .
```

Veya derleyip yürütülebilir oluşturun:

```powershell
go build ./...
./learn-go.exe   # Windows üzerinde
```

## Kısa Temel Açıklamalar (Basit ve Net)

- Paketler (`package`): Bir dizindeki Go dosyaları aynı `package`'i paylaşır. `package main` ve `func main()` programın başlangıç noktasıdır.
- Modüller (`go.mod`): Proje bağımlılıklarını ve modül adını tutar. Oluşturmak için `go mod init <module>` kullanılır.
- Tipler: `int`, `string`, `bool`, `rune`, `float64` gibi temel tipler. Her tipin bir "zero value" (varsayılan) değeri vardır.
- Değişkenler: `var x int = 5` veya fonksiyon içinde `x := 5` biçiminde tanımlanır.
- Sabitler: `const Pi = 3.14` veya `const ( A = iota )` ile blok halinde.

## Kontrol Akışları

- `if` / `else`: Koşula bağlı dallanma.
- `switch`: Birden çok durumu temiz şekilde ele alır; `switch` içinde değişken bildirimi yapılabilir.

Örnek: `if x := compute(); x > 0 { ... }`

## Fonksiyonlar ve Hata Yönetimi

- Go'da fonksiyonlar birden çok değer döndürebilir: `func f() (int, error)`.
- Hatalar tipik olarak `error` olarak döndürülür ve çağıran taraf `if err != nil {}` ile kontrol eder.

## Veri Yapıları: Diziler, Dilimler ve Haritalar

- Diziler sabit uzunluktadır: `var a [3]int`.
- Slice'lar dinamik, esnek yapıdır: `s := []int{1,2}` ve `s = append(s, 3)`.
- Map'ler anahtar-değer biçimindedir: `m := map[string]int{"a":1}`. Okurken `v, ok := m["a"]` ifadesi ile anahtarın varlığı kontrol edilir.

## Struct, Metot ve Interface

- `struct` ile birden fazla alanı gruplayabilirsiniz.
- Metot tanımlamak için `func (p Type) Method() {}` veya pointer receiver `func (p *Type) Mutate()` kullanılır.
- `interface` tipleri, metot kümelerini tanımlar; bir tip bu metotları sağlarsa interface'i uygulamış sayılır (implicit).

## İşaretçiler (Pointers)

- `&x` değişkenin adresini alır, `*p` işaretçiyi çözümler.
- Diziler değer tipi, slice ve map referans tipidir. Bu farklar kopyalama ve performans açısından önemlidir.

## Eşzamanlılık (Concurrency)

- `goroutine` ile fonksiyonları hafif thread'ler olarak çalıştırabilirsiniz: `go doWork()`.
- `channel`'lar goroutine'ler arası iletişim sağlar: `ch := make(chan int)`.
- `select` ile birden fazla kanalı dinleyebilirsiniz.

## Test Etme ve Araçlar

- Test dosyaları `_test.go` uzantısını alır ve `go test ./...` ile çalıştırılır.
- Kod formatlamak için `go fmt`, statik analiz için `go vet` veya `golangci-lint` kullanın.

## Hızlı Komutlar

```powershell
go run .
go build ./...
go test ./...
go fmt ./...
go mod tidy
```

## Katkıda Bulunma

- İsterseniz hata düzeltme, örnek ekleme veya açıklama değişiklikleri için PR açın.
- Yardımcı olmamı isterseniz PR başlığı ve açıklamasını hazırlayabilirim.

Her bölüm örnek olsun diye kod içinde zaten mevcut — dosyaları açıp değiştirip denemekten çekinmeyin. İsterseniz belirli bir bölüm için daha fazla örnek veya alıştırma ekleyeyim (ör. `context`, `error wrapping`, ileri concurrency desenleri, benchmark örnekleri).

İyi çalışmalar! 🎉
