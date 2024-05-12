package crypt

import "goms/maple"

type Crypter interface {
	Decrypt(buf []byte)
	Encrypt(buf []byte) []byte
}

type crypter struct {
	aes    *aesCrypt
	shanda Crypter
}

func NewCrypter() Crypter {
	var c crypter
	c.aes = NewAESCrypter()
	if maple.Region != 6 {
		c.shanda = newShandaCrypt()
	}
	return &c
}

// Decrypt implements Crypter.
func (c *crypter) Decrypt(buf []byte) {
	OFB(buf, c.aes.RecvIV16)
	if c.shanda != nil {
		c.shanda.Decrypt(buf)
	}
	Shuffle(&c.aes.RecvIV16)
}

// Encrypt implements Crypter.
func (c *crypter) Encrypt(data []byte) []byte {
	length := len(data)
	buf := make([]byte, length)
	_ = copy(buf, data)
	header := c.aes.EncodeHeader(length)
	if c.shanda != nil {
		_ = c.shanda.Encrypt(buf)
	}
	OFB(buf, c.aes.SendIV16)
	Shuffle(&c.aes.SendIV16)
	return append(header, buf...)
}
