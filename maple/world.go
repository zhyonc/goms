package maple

// World ID
type WorldID int8

// CMS World ID
const (
	BlueSnail WorldID = iota
	Shroom
	Slime
	RibbonPig
	JrNecki
	Lorang
	Tortie
	Octopus
	Lupin
	StarPixie
	JrPepe
	JrYeti
	StoneGolem
	WildKargo
	Hector
	WhiteBunny
	Manon
	FireBoar
	Ligator
	OrangeMushroom
)

var WorldNameMap map[WorldID]string = map[WorldID]string{
	BlueSnail:      "蓝蜗牛",
	Shroom:         "蘑菇仔",
	Slime:          "绿水灵",
	RibbonPig:      "漂漂猪",
	JrNecki:        "小青蛇",
	Lorang:         "红螃蟹",
	Tortie:         "大海龟",
	Octopus:        "章鱼怪",
	Lupin:          "顽皮猴",
	StarPixie:      "星精灵",
	JrPepe:         "胖企鹅",
	JrYeti:         "白雪人",
	StoneGolem:     "石头人",
	WildKargo:      "紫色猫",
	Hector:         "大灰狼",
	WhiteBunny:     "小白兔",
	Manon:          "喷火龙",
	FireBoar:       "火野猪",
	Ligator:        "青鳄鱼",
	OrangeMushroom: "花蘑菇",
}

// // TMS World ID
// const (
// 	Alia WorldID = iota
// 	Polytex
// 	Ryude
// 	Yuina
// 	Alicia
// 	KillerWhale
// 	Reboot WorldID = 45
// 	Tespia WorldID = 100
// )

// var NameMap map[WorldID]string = map[WorldID]string{
// 	Alia:        "艾麗亞",
// 	Polytex:     "普力特",
// 	Ryude:       "琉德",
// 	Yuina:       "優伊娜",
// 	Alicia:      "愛麗西亞",
// 	KillerWhale: "殺人鯨",
// 	Reboot:      "Reboot",
// 	Tespia:      "Tespia",
// }

type WorldTag uint8

const (
	FreeWorld WorldTag = iota
	ActivityWorld
	NewWorld
	FullWorld
)

type WorldStatus uint8

const (
	WorldIdle WorldStatus = iota
	WorldBusy
	WorldLimit
)

// World Ballon
type Ballon struct {
	NX      uint16
	NY      uint16
	Message string
}
