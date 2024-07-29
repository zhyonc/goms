package maple

const (
	NoneBurning  uint8 = 0
	BurningAcc   uint8 = 1
	MegaBurning  uint8 = 2
	LimitBurning uint8 = 3
)

const BurningMinLevel uint32 = 10

var BurningMaxLevel map[uint8]uint32 = map[uint8]uint32{
	NoneBurning:  0,
	BurningAcc:   130,
	MegaBurning:  150,
	LimitBurning: 200,
}
