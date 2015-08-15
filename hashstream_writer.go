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

func NewHashWriter(w io.Writer, h hash.Hash) (*HashWriter, error) {
	if w == nil || h == nil {
		return nil, errors.New("Reader or hash are nil: no allowed")
	}
	return &HashWriter{
		writer: w,
		hash:   h,
	}, nil
}

func NewMd5Writer(w io.Writer) *HashWriter {
	return &HashWriter{
		writer: w,
		hash:   md5.New(),
	}
}

func NewSha1Writer(w io.Writer) *HashWriter {
	return &HashWriter{
		writer: w,
		hash:   sha1.New(),
	}
}

func NewSha224Writer(w io.Writer) *HashWriter {
	return &HashWriter{
		writer: w,
		hash:   sha256.New224(),
	}
}

func NewSha256Writer(w io.Writer) *HashWriter {
	return &HashWriter{
		writer: w,
		hash:   sha256.New(),
	}
}

func NewSha384Writer(w io.Writer) *HashWriter {
	return &HashWriter{
		writer: w,
		hash:   sha512.New384(),
	}
}

func NewSha512Writer(w io.Writer) *HashWriter {
	return &HashWriter{
		writer: w,
		hash:   sha512.New(),
	}
}

type HashWriter struct {
	writer io.Writer
	hash   hash.Hash
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
