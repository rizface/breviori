package urlshortener_test

import (
	"testing"

	"github.com/rizface/breviori/urlshortener"
	"github.com/stretchr/testify/assert"
)

func TestKeyGen(t *testing.T) {
	lens := []int{8, 9, 10, 11}

	for _, l := range lens {
		key := urlshortener.KeyGen(l)
		assert.NotEqual(t, "", key)
		assert.True(t, (len(key) >= 8 && len(key) <= 11))
	}
}
