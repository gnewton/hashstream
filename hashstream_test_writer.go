package hashstream

import (
	"encoding/hex"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"testing"
)

func TestWriter(t *testing.T) {
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

	for {
		buf := make([]byte, bufSize)
		_, err := hr.Write(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			return "", err
		}
	}
	return hex.EncodeToString(hr.Sum()), nil

}
