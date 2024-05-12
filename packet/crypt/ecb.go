package crypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"goms/opcode"
	"math/rand"
	"strconv"
	"strings"
)

// Key
const (
	tmsKey string = "BrN=r54jQp2@yP6G" // maple version>=231
	gmsKey string = "N3x@nGLEUH@ckEr!"
)

var (
	secretKey      string        = tmsKey
	RandNumStrLen  int           = 4
	EncryptOpcodes map[int]int16 = make(map[int]int16)
	EncryptContent []byte
)

func init() {
	// Use random numbers to represent opcodes
	min := opcode.BeginUser // 203
	max := opcode.EndUser   // 2048
	var builder strings.Builder
	for i := min; i < max; i++ {
		randNum := rand.Intn(9999-int(min)+1) + int(min)
		for {
			// Need unique rand num
			_, ok := EncryptOpcodes[randNum]
			if ok {
				// Already contain the rand num
				randNum = rand.Intn(9999-int(min)+1) + int(min)
				continue
			}
			break
		}
		EncryptOpcodes[randNum] = int16(i)
		randNumStr := fmt.Sprintf("%0"+strconv.Itoa(RandNumStrLen)+"d", randNum) // 240->"0240"
		builder.WriteString(randNumStr)
	}
	content := []byte(builder.String())
	// Des key
	keyBytes := make([]byte, 24)
	copy(keyBytes, []byte(secretKey))
	copy(keyBytes[16:], keyBytes[:8]) // BrN=r54jQp2@yP6G->BrN=r54jQp2@yP6GBrN=r54j
	block, err := des.NewTripleDESCipher(keyBytes)
	if err != nil {
		panic(err)
	}
	// New cipher.BlockMode
	blockMode := newECB(block)
	// Pad plaintext to an integer multiple of the block size
	padding := block.BlockSize() - len(content)%block.BlockSize() // 4
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)       // PKCS5/PKCS7
	content = append(content, padtext...)                         // []byte{4,4,4,4}
	// Done
	EncryptContent = make([]byte, len(content))
	blockMode.Encrypt(EncryptContent, content)
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func (x *ecb) BlockSize() int {
	return x.blockSize
}

func (x *ecb) Encrypt(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func (x *ecb) Decrypt(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
