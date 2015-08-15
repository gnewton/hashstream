// package provides Writer/Reader support for large files or streaming, inspired by Java's DigestInputStream/DigestOutputStream https://docs.oracle.com/javase/7/docs/api/java/security/DigestInputStream.html
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

func NewMD5Reader(r io.Reader) *HashReader {
	return &HashReader{
		reader: r,
		hash:   md5.New(),
	}
}

func NewSHA1Reader(r io.Reader) *HashReader {
	return &HashReader{
		reader: r,
		hash:   sha1.New(),
	}
}

func NewSHA224Reader(r io.Reader) *HashReader {
	return &HashReader{
		reader: r,
		hash:   sha256.New224(),
	}
}

func NewSHA256Reader(r io.Reader) *HashReader {
	return &HashReader{
		reader: r,
		hash:   sha256.New(),
	}
}

func NewSHA384Reader(r io.Reader) *HashReader {
	return &HashReader{
		reader: r,
		hash:   sha512.New384(),
	}
}

func NewSHA512Reader(r io.Reader) *HashReader {
	return &HashReader{
		reader: r,
		hash:   sha512.New(),
	}
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
