package maple

type Local uint8

const (
	CN     Local = 4
	TW     Local = 6
	Global Local = 8
)

const (
	Version      uint16 = 138
	MinorVersion string = "1"
	Region       Local  = CN
	NexonIP      string = "221.231.130.70" // Use ijl15.dll forward this ip to private ip
)

const (
	MaxCharacterSlot       uint8 = 40
	MaxQuickSlot           uint8 = 32
	FamiliarCardNameLength int   = 11
	CharacterNameLength    int   = 13
)
