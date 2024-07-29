package crypter

import "goms/maple"

type Crypter interface {
	Decrypt(buf []byte)
	Encrypt(buf []byte) []byte
	SetTempRecvIV(buf []byte)
}

type crypter struct {
	aes    *aesCrypter
	shanda *shandaCrypter
}

func NewCrypter(recvIV, sendIV [4]byte) Crypter {
	var c crypter
	c.aes = NewAESCrypter(recvIV, sendIV)
	if maple.Region == maple.CN && maple.Version < 119 {
		// for very old version
		c.shanda = newShandaCrypter()
	}
	return &c
}

// SetTempRecvIV implements Crypter.
func (c *crypter) SetTempRecvIV(newIV []byte) {
	for i := 0; i < 4; i++ {
		copy(c.aes.RecvIV16[4*i:], newIV)
	}
	Shuffle(&c.aes.RecvIV16)
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
