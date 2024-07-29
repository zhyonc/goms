package maple

type EquipPart uint16

const (
	_ EquipPart = iota
	Hat
	FaceAcc
	EyeAcc
	EarAcc
	Top
	Bottom
	Shoes
	Gloves
	Cape
	SubWeapon
	Weapon
	Ring1
	Ring2
	Pet1Acc
	Ring3
	Ring4
	Pendant
	TamingMob // 骑宠
	Saddle    // 鞍
	MobEquip
	Medal
	Belt
	Shoulder
	Pet2Acc
	Pet3Acc
	Pocket
	_ // Old Android
	_ // Old AndroidHeart
	Badge
	Emblem
	Extended0
	Extended1
	Extended2
	Extended3
	Extended4
	Extended5
	Extended6
	Overall                       = Top // Top and overall share the same body part
	Android             EquipPart = 53
	AndroidHeart        EquipPart = 54
	MonsterBook         EquipPart = 55
	PendantExt          EquipPart = 65
	CashWeapon          EquipPart = 100
	PetConsumeHPItem    EquipPart = 200
	PetConsumeMPItem    EquipPart = 201
	EvanHat             EquipPart = 1000
	EvanPendant         EquipPart = 1001
	EvanWing            EquipPart = 1002
	EvanShoes           EquipPart = 1003
	MechineEngine       EquipPart = 1100
	MachineArm          EquipPart = 1101
	MachineLeg          EquipPart = 1102
	MachineFrame        EquipPart = 1103
	MachineTransistor   EquipPart = 1104
	APHat               EquipPart = 1200
	APCape              EquipPart = 1201
	APFaceAccessory     EquipPart = 1202
	APTop               EquipPart = 1203
	APOverall           EquipPart = APTop
	APBottom            EquipPart = 1204
	APShoes             EquipPart = 1205
	APGloves            EquipPart = 1206
	DUHat               EquipPart = 1300
	DUCape              EquipPart = 1301
	DUFaceAccessory     EquipPart = 1302
	DUGloves            EquipPart = 1304
	DUEyeAccessory      EquipPart = 1305
	DUEarrings          EquipPart = 1306
	DUTop               EquipPart = 1307
	DUOverall           EquipPart = DUTop
	DUBottom            EquipPart = 1308
	DUShoes             EquipPart = 1309
	BitsBase            EquipPart = 1400 // 1400~1424
	BitsEnd             EquipPart = 1425
	ZeroEyeAccessory    EquipPart = 1500
	ZeroHat             EquipPart = 1501
	ZeroFaceAccessory   EquipPart = 1502
	ZeroEarrings        EquipPart = 1503
	ZeroCape            EquipPart = 1504
	ZeroTop             EquipPart = 1505
	ZeroOverall         EquipPart = 1505
	ZeroGloves          EquipPart = 1506
	ZeroWeapon          EquipPart = 1507
	ZeroBottom          EquipPart = 1508
	ZeroShoes           EquipPart = 1509
	ZeroRing1           EquipPart = 1510
	ZeroRing2           EquipPart = 1511
	ZeroPendant         EquipPart = 1512 // ?
	ZeroPendantExt      EquipPart = 1513 // ?
	AFVanishinJourney   EquipPart = 1600
	AFChuChu            EquipPart = 1601
	AFLachelein         EquipPart = 1602
	AFArcana            EquipPart = 1603
	AFMorass            EquipPart = 1604
	AFEsfera            EquipPart = 1605
	Totem1              EquipPart = 5000
	Totem2              EquipPart = 5001
	Totem3              EquipPart = 5002
	MBPHat              EquipPart = 5101
	MBPCape             EquipPart = 5102
	MBPTop              EquipPart = 5103
	MBPOverall          EquipPart = 5103
	MBPGloves           EquipPart = 5104
	MBPShoes            EquipPart = 5105
	MBPWeapon           EquipPart = 5106
	HakuFan             EquipPart = 5200
	SlotIndexNotDefined EquipPart = 15440
	VEIBase             EquipPart = 20000
	VEIEnd              EquipPart = 20024
)

type EquipStat uint32

const ( // GW_ItemSlotEquipBase::Decode
	RUC              EquipStat = 0x1
	CUC              EquipStat = 0x2
	STR              EquipStat = 0x4
	DEX              EquipStat = 0x8
	INT              EquipStat = 0x10
	LUK              EquipStat = 0x20
	MaxHP            EquipStat = 0x40
	MaxMP            EquipStat = 0x80
	PAD              EquipStat = 0x100
	MAD              EquipStat = 0x200
	PDD              EquipStat = 0x400
	MDD              EquipStat = 0x800
	ACC              EquipStat = 0x1000
	EVA              EquipStat = 0x2000
	Craft            EquipStat = 0x4000
	Speed            EquipStat = 0x8000
	Jump             EquipStat = 0x10000
	Attribute        EquipStat = 0x20000
	LevelUpType      EquipStat = 0x40000
	Level            EquipStat = 0x80000
	EXP              EquipStat = 0x100000
	Durability       EquipStat = 0x200000
	IUC              EquipStat = 0x400000
	PVPDamage        EquipStat = 0x800000
	ReduceReq        EquipStat = 0x1000000
	SpecialAttribute EquipStat = 0x2000000
	DurabilityMax    EquipStat = 0x4000000
	IncReq           EquipStat = 0x8000000
	GrowthEnchant    EquipStat = 0x10000000
	PSEnchant        EquipStat = 0x20000000
	BDR              EquipStat = 0x40000000
	MDR              EquipStat = 0x80000000 // >=0？
	DamR             EquipStat = 0x1
	StatR            EquipStat = 0x2
	Cuttable         EquipStat = 0x4
	ExGradeOption    EquipStat = 0x8
	ItemState        EquipStat = 0x10
)
