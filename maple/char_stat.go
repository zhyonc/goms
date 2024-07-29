package maple

type CharStat uint64

const (
	CS_None          CharStat = 0
	CS_Skin          CharStat = 0x1
	CS_Face          CharStat = 0x2
	CS_Hair          CharStat = 0x4
	CS_Level         CharStat = 0x10
	CS_Job           CharStat = 0x20
	CS_STR           CharStat = 0x40
	CS_DEX           CharStat = 0x80
	CS_INT           CharStat = 0x100
	CS_LUK           CharStat = 0x200
	CS_HP            CharStat = 0x400
	CS_MHP           CharStat = 0x800
	CS_MP            CharStat = 0x1000
	CS_MMP           CharStat = 0x2000
	CS_AP            CharStat = 0x4000
	CS_ExtendSP      CharStat = 0x8000
	CS_EXP           CharStat = 0x10000
	CS_POP           CharStat = 0x20000
	CS_Money         CharStat = 0x40000
	CS_Fatigue       CharStat = 0x80000
	CS_CharismaEXP   CharStat = 0x100000
	CS_InsightEXP    CharStat = 0x200000
	CS_WillEXP       CharStat = 0x400000
	CS_CraftEXP      CharStat = 0x800000
	CS_SenseEXP      CharStat = 0x1000000
	CS_CharmEXP      CharStat = 0x2000000
	CS_DayLimit      CharStat = 0x4000000
	CS_AlbaActivity  CharStat = 0x8000000
	CS_CharacterCard CharStat = 0x10000000
	CS_PVP1          CharStat = 0x20000000 // EXP and Grade
	CS_PVP2          CharStat = 0x40000000 // ModeLevel and ModeType
	CS_EventPoint    CharStat = 0x80000000
	CS_All           CharStat = 0xFFFFFFFF
)
