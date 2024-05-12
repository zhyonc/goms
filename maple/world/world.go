package world

// World ID
type WorldID uint8

const (
	Alia WorldID = iota
	Polytex
	Ryude
	Yuina
	Alicia
	KillerWhale
	Reboot WorldID = 45
	Tespia WorldID = 100
)

// World Name
var NameMap map[WorldID]string = map[WorldID]string{
	Alia:        "艾麗亞",
	Polytex:     "普力特",
	Ryude:       "琉德",
	Yuina:       "優伊娜",
	Alicia:      "愛麗西亞",
	KillerWhale: "殺人鯨",
	Reboot:      "Reboot",
	Tespia:      "Tespia",
}

// World Event
type WorldEvent uint8

const (
	Burning  WorldEvent = 0
	PinkBean WorldEvent = 1
	Normal   WorldEvent = 255
)

// World State
type WorldState uint8

const (
	Free WorldState = iota
	Activity
	New
	Full
)

// World Ballon
type Ballon struct {
	NX      uint16
	NY      uint16
	Message string
}
