package main

import (
	"crypto/md5"
	"hash"
	"hash/crc32"
)

func NewCRCHasher() *Hasher {

	return &Hasher{
		h: crc32.NewIEEE(),
	}

}

func NewMD5Hasher() *Hasher {

	return &Hasher{
		h: md5.New(),
	}

}

type Hasher struct {
	h hash.Hash
}

func (c *Hasher) Write(p []byte) (n int, err error) {
	return c.h.Write(p)
}

func (c *Hasher) Sum() []byte {
	return c.h.Sum(nil)
}
