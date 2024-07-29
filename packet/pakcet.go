package packet

import (
	"goms/maple"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

var (
	LgDecoder *encoding.Decoder
	LgEncoder *encoding.Encoder
)

func init() {
	if maple.Region == maple.CN {
		LgDecoder = simplifiedchinese.GBK.NewDecoder()
		LgEncoder = simplifiedchinese.GBK.NewEncoder()
	} else if maple.Region == maple.TW {
		LgDecoder = traditionalchinese.Big5.NewDecoder()
		LgEncoder = traditionalchinese.Big5.NewEncoder()
	} else {
		LgDecoder = encoding.Nop.NewDecoder()
		LgEncoder = encoding.Nop.NewEncoder()
	}
}
