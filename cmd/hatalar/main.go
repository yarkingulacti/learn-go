package hatalar

import (
	"fmt"
	"log"
)

type farkSayHatasi struct {
	kaynak string
	hedef  string
}

func (d farkSayHatasi) Error() string {
	return fmt.Sprintf("Fark Say Hatası: %s kaynağının %s hedefinden bir uzunluk farkı yok!", d.kaynak, d.hedef)
}

func farkSay(kaynak string, hedef string) (int, error) {
	if len(kaynak) == len(hedef) {
		return 0, farkSayHatasi{kaynak: kaynak, hedef: hedef}
	}

	if len(kaynak) > len(hedef) {
		return len(kaynak) - len(hedef), nil
	}

	return len(hedef) - len(kaynak), nil
}

func bolme(a, b int) (float64, error) {
	if b == 0 {
		// panic(errors.New("Sıfıra bölme hatası!"))
		log.Fatal("Sıfıra bölme hatası")
	}

	return float64(a) / float64(b), nil
}

func HatalarIcerik() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panik modundan çıkıldı! Alınan hata: %s\n", e)
		}
	}()
	fark, hata := farkSay("abcde", "abcd")

	if hata != nil {
		fmt.Println("Hata oluştu:", hata)
	} else {
		fmt.Println("Fark sayısı:", fark)
	}

	bolme(10, 0)
}

func main() {
	HatalarIcerik()
}
