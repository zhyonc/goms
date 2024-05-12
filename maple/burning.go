package maple

type BurningType uint8

const (
	NoneBurning  BurningType = 0
	BurningAcc   BurningType = 1
	MegaBurning  BurningType = 2
	LimitBurning BurningType = 3
)

const BurningMinLevel uint32 = 10

var BurningMaxLevel map[BurningType]uint32 = map[BurningType]uint32{
	NoneBurning:  0,
	BurningAcc:   130,
	MegaBurning:  150,
	LimitBurning: 200,
}
