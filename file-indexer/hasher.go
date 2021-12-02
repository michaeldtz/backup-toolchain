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

func NewHasher() *Hasher {

	return &Hasher{
		h: md5.New(),
	}

}

type Hasher struct {
	h hash.Hash
}

func (c *Hasher) Write(p []byte) (n int, err error) {
	c.h.Write(p)
	return
}

func (c *Hasher) Sum() []byte {
	return c.h.Sum(nil)
}
