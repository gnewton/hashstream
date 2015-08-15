// package provides Writer/Reader support for large files or streaming, inspired by Java's DigestInputStream/DigestOutputStream https://docs.oracle.com/javase/7/docs/api/java/security/DigestInputStream.html
package hashstream

import (
	"crypto/md5"
	"errors"
	"hash"
	"io"
)

func NewHashReader(r io.Reader, h hash.Hash) *HashReader {
	return &HashReader{
		reader: r,
		hash:   h,
	}
}

func NewMD5Reader(r io.Reader) *HashReader {
	return &HashReader{
		reader: r,
		hash:   md5.New(),
	}
}

func NewHashWriter(w io.Writer, h hash.Hash) *HashWriter {
	return &HashWriter{
		writer: w,
		hash:   h,
	}
}

func NewMD5Writer(w io.Writer) *HashWriter {
	return &HashWriter{
		writer: w,
		hash:   md5.New(),
	}
}

type HashReader struct {
	reader io.Reader
	hash   hash.Hash
}

type HashWriter struct {
	writer io.Writer
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

func (l *HashWriter) Write(p []byte) (n int, err error) {
	n, err = l.writer.Write(p)
	if err != nil {
		return -1, err
	}
	l.hash.Write(p)

	return
}

func (l *HashWriter) Sum() []byte {
	return sum(l.hash)
}
