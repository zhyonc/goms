package util

func SimpleXOR(buf []byte, keyBytes []byte) {
	keyLength := len(keyBytes)
	for i, b := range buf {
		buf[i] = b ^ keyBytes[i%keyLength]
	}
}
