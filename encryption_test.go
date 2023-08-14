package hibbettencryptiongo_test

import (
	"log"
	"testing"

	hibbettencryptiongo "github.com/KakashiHatake324/hibbett-encryption-go"
)

func TestEncrypt(t *testing.T) {
	out, err := hibbettencryptiongo.EncryptCard("4242424242424242")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(out)
}
