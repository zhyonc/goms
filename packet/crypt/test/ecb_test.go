package test

import (
	"fmt"
	"goms/packet/crypt"
	"testing"
)

func TestOpcodeEncryption(t *testing.T) {
	fmt.Println(crypt.EncryptOpcodes)
	fmt.Println(crypt.EncryptContent)
}
