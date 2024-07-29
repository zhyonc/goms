package test

import (
	"goms/opcode"
	"testing"
)

const (
	inPath     string = "../in.go"
	inGenPath  string = "../in_gen.go"
	inMapName  string = "InMap"
	outPath    string = "../out.go"
	outGenPath string = "../out_gen.go"
	outMapName string = "OutMap"
)

func TestGenOpcodeMap(t *testing.T) {
	opcode.GenOpcodeMap(inPath, inGenPath, inMapName)
	opcode.GenOpcodeMap(outPath, outGenPath, outMapName)
}
