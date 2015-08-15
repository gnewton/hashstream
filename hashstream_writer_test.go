package hashstream

import (
	"crypto/md5"
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

func TestWriter_NilWriter(t *testing.T) {
	_, err := NewHashWriter(nil, md5.New())
	if err == nil {
		t.FailNow()
	}
}

func TestWriter_NilHash(t *testing.T) {
	writer := ioutil.Discard
	_, err := NewHashWriter(writer, nil)
	if err == nil {
		t.FailNow()
	}
}

func applyWriterHash(text string, hs hash.Hash, bufSize int) (string, error) {
	writer := ioutil.Discard
	hr, err := NewHashWriter(writer, hs)
	if err != nil {
		return "", nil
	}

	_, err = hr.Write([]byte(text))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return hex.EncodeToString(hr.Sum()), nil

}
