package maple

type ForcedStat uint32

const (
	FS_None     ForcedStat = 0
	FS_STR      ForcedStat = 0x1
	FS_DEX      ForcedStat = 0x2
	FS_INT      ForcedStat = 0x4
	FS_LUK      ForcedStat = 0x8
	FS_PAD      ForcedStat = 0x10
	FS_MAD      ForcedStat = 0x20
	FS_PDD      ForcedStat = 0x40
	FS_MDD      ForcedStat = 0x80
	FS_ACC      ForcedStat = 0x100
	FS_EVA      ForcedStat = 0x200
	FS_Speed    ForcedStat = 0x400
	FS_Jump     ForcedStat = 0x800
	FS_SpeedMax ForcedStat = 0x1000
	FS_OptOff   ForcedStat = 0x2000
	FS_AddMHP   ForcedStat = 0x4000
)
