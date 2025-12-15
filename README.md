% Learn Go

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
- **Yerleşik paralellik desteği**: `goroutine` ve `channel` ile hafif paralel işlemler.
- **Basit sözdizimi**: Okunabilir ve öğrenmesi kolay.
- **Standart araç zinciri**: `go fmt`, `go vet`, `go test`, `go build`, `go mod` vb.

---

## 2. Kurulum ve Araçlar

1. Go'nun resmi sitesinden uygun sürümü indirin: https://go.dev/dl/
2. Ortam değişkenleri (`GOROOT`, `GOPATH`) genellikle otomatik ayarlanır; mod tabanlı projelerde `GOPATH` artık kritik değildir.

Yaygın `go` komutları:

- `go run main.go` — Kaynak dosyayı derleyip çalıştırır.
- `go build` — Paketleri derler, yürütülebilir dosya üretir.
- `go test` — Testleri çalıştırır.
- `go fmt` — Kodunuzu biçimlendirir.
- `go vet` — Potansiyel hataları bulmaya çalışır.
- `go mod init <module>` — Modül oluşturur ve `go.mod` dosyası ekler.
- `go get` — Bağımlılıkları yönetir (modlarla birlikte `go get` davranışı değişti).

Örnek: Proje başlatmak

```powershell
go mod init github.com/username/projectname
go fmt ./...
go build ./...
```

---

## 3. Modüller ve Paketler

Go kodu paketlere (`package`) ayrılır. Her dizin bir paket barındırır (istisnalar var). Projeler `go.mod` dosyası ile modül haline getirilir; bu dosya bağımlılıkları ve modül adı bilgisini tutar.

`main` paketi bir yürütülebilir program üretir ve `main()` fonksiyonu programın giriş noktasıdır.

Örnek paket yapısı:

```
project/
  go.mod
  main.go        // package main
  greet/
  greet.go     // package greet
```

`greet/greet.go`:

```go
package greet

func Hello(name string) string {
  return "Hello, " + name
}
```

`main.go`:

```go
package main

import (
  "fmt"
  "project/greet"
)

func main() {
  fmt.Println(greet.Hello("Dünya"))
}
```

---

## 4. Temel Tipler

Go'da temel tipler:

- Sayısal tipler: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8` (`byte`), `float32`, `float64`, `complex64`, `complex128`.
- Metin: `string` (UTF-8 olarak saklanır).
- Karakter/rune: `rune` (alias `int32`) Unicode kod noktasıdır.
- Boolean: `bool`.

Her tipin sıfır değeri (zero value) vardır: sayılar için `0`, string için `""`, bool için `false`, pointer, slice, map, channel, func ve interface için `nil`.

Örnek:

```go
var i int        // 0
var s string     // ""
var b bool       // false
var r rune       // 0
```

---

## 5. Değişkenler ve Sabitler

Değişken bildirimleri:

- Paket-scope: `var x int = 5`
- Kısa-değişken atama (sadece fonksiyon içinde): `x := 5`
- Çoklu atama: `var a, b = 1, "two"` veya `a, b := 1, "two"`

Sabitler:

```go
const Pi = 3.14159
const (
  A = iota // 0
  B = iota // 1
)
```

`iota` sabit bloklarında otomatik artan tamsayı üretmek için kullanılır.

---

## 6. Kontrol Akışları

For döngüsü (Go'da tek döngü yapısı `for`):

```go
for i := 0; i < 10; i++ {
  // ...
}

for index, value := range slice {
  // ...
}
```

If/else:

```go
if x := compute(); x > 0 {
  // ...
} else {
  // ...
}
```

Switch (etiket gerektirmez):

```go
switch v := interfaceVal.(type) {
case int:
  // tiplendirilmiş işlem
case string:
  // ...
default:
  // ...
}
```

---

## 7. Fonksiyonlar

Fonksiyon tanımı ve özellikleri:

```go
func add(a int, b int) int {
  return a + b
}

// Çoklu dönüş
func swap(a, b string) (string, string) {
  return b, a
}

// Variadic fonksiyon
func sum(nums ...int) int {
  total := 0
  for _, n := range nums {
    total += n
  }
  return total
}

// Closure
func makeAdder(x int) func(int) int {
  return func(y int) int { return x + y }
}
```

Fonksiyonlar birinci sınıf vatandaştır; fonksiyonu parametre olarak geçirmek veya dönüş değeri olarak kullanmak mümkündür.

---

## 8. Diziler, Dilimler (Slices) ve Haritalar (Maps)

Diziler sabit uzunlukludur:

```go
var a [3]int = [3]int{1,2,3}
```

Slice, dinamik ve daha çok kullanılır:

```go
s := []int{1,2,3}
s = append(s, 4)
sub := s[1:3] // dilimleme
```

Map (hash tabanlı sözlük):

```go
m := map[string]int{"a": 1}
val, ok := m["a"]
if !ok {
  // anahtar yok
}
delete(m, "a")
```

Not: Slices ve maps referans tiplerdir; kopyalandıklarında veri paylaşımı olabilir.

---

## 9. Struct'lar, Metodlar ve Interface'ler

Struct tanımı:

```go
type Person struct {
  Name string
  Age  int
}

func (p Person) Greet() string {
  return "Hi, I'm " + p.Name
}

// Pointer receiver
func (p *Person) HaveBirthday() {
  p.Age++
}
```

Interface örneği:

```go
type Greeter interface {
  Greet() string
}

func SayHello(g Greeter) {
  fmt.Println(g.Greet())
}
```

Go'da interface'ler implicit (zımni) olarak uygulanır: bir tip interface'in tüm metotlarını sağlıyorsa o interface'i uygulamış sayılır.

---

## 10. İşaretçiler (Pointers)

Go'da pointer kullanımı:

```go
var x int = 10
var p *int = &x
*p = 20 // x artık 20
```

Pointerlar ile değişkenin kendisini değiştirebilir ve büyük verilerin kopyalanmasını önleyebilirsiniz.

---

## 11. Hata Yönetimi

Go'da hata yönetimi tipik olarak dönüş değeri olarak `error` döndürülerek yapılır.

```go
import "errors"

func doSomething() error {
  if fail {
    return errors.New("bir hata oluştu")
  }
  return nil
}

if err := doSomething(); err != nil {
  // hata işleme
}
```

Go 1.13 ile gelen hata sarmalama (`fmt.Errorf("", %w)`) sayesinde hata zinciri izlenebilir.

---

## 12. Eşzamanlılık: Goroutine'ler ve Channel'lar

Go'nun güçlü olduğu alanlardan biri eşzamanlılıktır.

Goroutine başlatma:

```go
go doWork() // arka planda çalışır
```

Channel:

```go
ch := make(chan int)
go func() { ch <- 42 }()
val := <-ch
```

Buffered channel:

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
```

`select` ile birden çok kanal arasında seçim:

```go
select {
case v := <-ch1:
  fmt.Println("ch1", v)
case ch2 <- 5:
  fmt.Println("yazıldı")
default:
  fmt.Println("hiçbiri hazır değil")
}
```

Eşzamanlı programlama için ayrıca `sync.WaitGroup`, `sync.Mutex` gibi senkronizasyon primitifleri mevcuttur.

---

## 13. Test Etme

Go, `testing` paketini içerir. Test dosyaları `_test.go` uzantısı alır.

```go
import "testing"

func TestAdd(t *testing.T) {
  if add(2,3) != 5 {
    t.Fatal("beklenmeyen sonuç")
  }
}
```

Çalıştırma:

```bash
go test ./...
```

---

## 14. Kod Kalitesi ve Araçlar

- `go fmt` ile formatlama
- `golangci-lint` veya `go vet` ile statik analiz
- `go mod tidy` ile kullanılmayan modülleri temizleme

Örnek komutlar:

```bash
go fmt ./...
go vet ./...
go test ./...
go mod tidy
```

---

## 15. Performans ve Profiling

Go, `pprof` ve `trace` gibi yerleşik araçlarla CPU ve bellek profil çıkarma desteği sağlar. Büyük hizmetlerde performans analizi yapmak için bu araçlar önemlidir.

---

## 16. En İyi Uygulamalar (Best Practices)

- Basit tutun; karmaşık çözümlerden kaçının.
- İsimlendirme: kısa ve anlamlı isimler kullanın.
- Konuşlandırılabilir paketleri küçük tutun.
- Hataları hemen kontrol edin (`if err != nil { return err }`).
- `context.Context` kullanarak iptal ve zaman aşımı pas geçmeyin.
- Test yazın ve CI'de çalıştırın.

---

## 17. Kaynaklar

- Resmi dokümantasyon: https://go.dev/
- Effective Go: https://go.dev/doc/effective_go
- Tour of Go: https://tour.golang.org/

---

Bu rehberi geliştirmek veya belirli bir konu için örnekler eklememi isterseniz belirtin; örneğin `context` kullanımı, paket düzeni, ileri seviye concurrency desenleri veya test ve benchmark örnekleri ekleyebilirim.
