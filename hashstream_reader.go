// Provides crypto hash Writer/Reader support for large files or streaming.
//
// Inspired by Java's DigestInputStream/DigestOutputStream https://docs.oracle.com/javase/7/docs/api/java/security/DigestInputStream.html
//
// Example: reader
//
//	package main
//
//	import (
//		"crypto/md5"
//		"encoding/hex"
//		"fmt"
//		"github.com/gnewton/hashstream"
//		"io"
//		"log"
//		"strings"
//	)
//
//	func main() {
//		reader := strings.NewReader("hello")
//		hr, err := hashstream.NewHashReader(reader, md5.New())
//		for {
//			buf := make([]byte, 4)
//			_, err = hr.Read(buf)
//			if err != nil {
//				if err == io.EOF {
//					break
//				}
//				log.Fatal(err)
//			}
//		}
//		fmt.Println(hex.EncodeToString(hr.Sum()))
//	}
//
//      Output: 5d41402abc4b2a76b9719d911017c592
//
// Example: writer
//
package hashstream

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"io"
)

func NewHashReader(r io.Reader, h hash.Hash) (*HashReader, error) {
	if r == nil || h == nil {
		return nil, errors.New("Reader or hash are nil: no allowed")
	}
	return &HashReader{
		reader: r,
		hash:   h,
	}, nil
}

func NewMD5Reader(r io.Reader) (*HashReader, error) {
	return NewHashReader(r, md5.New())
}

func NewSHA1Reader(r io.Reader) (*HashReader, error) {
	return NewHashReader(r, sha1.New())

}

func NewSHA224Reader(r io.Reader) (*HashReader, error) {
	return NewHashReader(r, sha256.New224())
}

func NewSHA256Reader(r io.Reader) (*HashReader, error) {
	return NewHashReader(r, sha256.New())
}

func NewSHA384Reader(r io.Reader) (*HashReader, error) {
	return NewHashReader(r, sha512.New384())
}

func NewSHA512Reader(r io.Reader) (*HashReader, error) {
	return NewHashReader(r, sha512.New())
}

type HashReader struct {
	reader io.Reader
	hash   hash.Hash
}

func (l *HashReader) Sum() []byte {
	return sum(l.hash)
}

func sum(h hash.Hash) []byte {
	return h.Sum(nil)
}

func (l *HashReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		n = -1
		err = errors.New("Buffer cannot be length=zero")
		return
	}

	n, err = l.reader.Read(p)

	if err != nil {
		return -1, err
	}
	if n > 0 {
		var buf []byte
		if n == len(p) {
			buf = p
		} else {
			buf = p[0:n]
		}
		n, err = l.hash.Write(buf)
		if err != nil {
			return -1, nil
		}
	}
	return
}
