package hashstream

import (
	"encoding/hex"
	"hash"
	"io/ioutil"
	"log"
	"testing"
)

func TestWriter(t *testing.T) {
	resetCrypto()
	for text, tcrm := range textCryptoResultMap {
		for hsh, sum := range tcrm {
			hexhash, err := applyWriterHash(text, hsh, 32)
			if err != nil || hexhash != sum {
				t.FailNow()
			}
		}
	}

}

func applyWriterHash(text string, hs hash.Hash, bufSize int) (string, error) {
	writer := ioutil.Discard
	hr := NewHashWriter(writer, hs)

	_, err := hr.Write([]byte(text))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return hex.EncodeToString(hr.Sum()), nil

}
