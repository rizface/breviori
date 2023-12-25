package urlshortener

import (
	"math/rand"
)

func KeyGen(keyLen int) string {
	key := ""
	alphanumchars := []rune("012@34[567]89AB^CDEF#GHI!JKL%MN^OP@QRST^UV@WX!YZ@abc$de#fghijklmnopqrstuvwxyz")

	for i := 0; i < keyLen; i++ {
		key += string(alphanumchars[rand.Intn(len(alphanumchars))])
	}

	return key
}
