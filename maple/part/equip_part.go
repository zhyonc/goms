package part

var EquipPartMap map[string]uint8 = map[string]uint8{
	"Hat":          Hat,
	"FaceAcc":      FaceAcc,
	"EyeAcc":       EyeAcc,
	"EarAcc":       EarAcc,
	"Top":          Top,
	"Bottom":       Bottom,
	"Overall":      Overall,
	"Shoes":        Shoes,
	"Gloves":       Gloves,
	"Cape":         Cape,
	"SubWeapon":    SubWeapon,
	"Weapon":       Weapon,
	"Ring1":        Ring1,
	"Ring2":        Ring2,
	"Ring3":        Ring3,
	"Ring4":        Ring4,
	"Pendant":      Pendant,
	"TamingMob":    TamingMob,
	"Saddle":       Saddle,
	"MobEquip":     MobEquip,
	"Medal":        Medal,
	"Belt":         Belt,
	"Shoulder":     Shoulder,
	"Pocket":       Pocket,
	"Badge":        Badge,
	"Emblem":       Emblem,
	"Android":      Android,
	"AndroidHeart": AndroidHeart,
}

const (
	_ uint8 = iota
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
	TamingMob // 騎寵
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
	Overall            = Top // Top and overall share the same body part
	Android      uint8 = 53
	AndroidHeart uint8 = 54
	MonsterBook  uint8 = 55
	PendantExt   uint8 = 65
)
