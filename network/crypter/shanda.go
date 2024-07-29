package crypter

// ShandaCrypto (which became non-used after CMSV118 or GMSv149)
// I don't sure if this crypter works
type shandaCrypter struct {
}

func newShandaCrypter() *shandaCrypter {
	return &shandaCrypter{}
}

// Taken from Kagami
func (sc *shandaCrypter) Decrypt(buf []byte) {
	var j int32
	var a, b, c byte

	for i := byte(0); i < 3; i++ {
		a = 0
		b = 0

		for j = int32(len(buf)); j > 0; j-- {
			c = buf[j-1]
			c = rol(c, 3)
			c ^= 0x13
			a = c
			c ^= b
			c = byte(int32(c) - j)
			c = ror(c, 4)
			b = a
			buf[j-1] = c
		}

		a = 0
		b = 0

		for j = int32(len(buf)); j > 0; j-- {
			c = buf[int32(len(buf))-j]
			c -= 0x48
			c ^= 0xFF
			c = rol(c, int(j))
			a = c
			c ^= b
			c = byte(int32(c) - j)
			c = ror(c, 3)
			b = a
			buf[int32(len(buf))-j] = c
		}
	}
}

// Taken from Kagami
func (sc *shandaCrypter) Encrypt(buf []byte) []byte {
	var j int32
	var a, c byte
	for i := byte(0); i < 3; i++ {
		a = 0

		for j = int32(len(buf)); j > 0; j-- {
			c = buf[int32(len(buf))-j]
			c = rol(c, 3)
			c = byte(int32(c) + j)
			c ^= a
			a = c
			c = ror(a, int(j))
			c ^= 0xFF
			c += 0x48
			buf[int32(len(buf))-j] = c
		}

		a = 0

		for j = int32(len(buf)); j > 0; j-- {
			c = buf[j-1]
			c = rol(c, 4)
			c = byte(int32(c) + j)
			c ^= a
			a = c
			c ^= 0x13
			c = ror(c, 3)
			buf[j-1] = c
		}
	}
	return nil
}

// Taken from Kagami
func ror(val byte, num int) byte {
	for i := 0; i < num; i++ {
		var lowbit int

		if val&1 > 0 {
			lowbit = 1
		} else {
			lowbit = 0
		}

		val >>= 1
		val |= byte(lowbit << 7)
	}

	return val
}

// Taken from Kagami
func rol(val byte, num int) byte {
	var highbit int

	for i := 0; i < num; i++ {
		if val&0x80 > 0 {
			highbit = 1
		} else {
			highbit = 0
		}

		val <<= 1
		val |= byte(highbit)
	}

	return val
}
