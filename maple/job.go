package maple

const (
	Beginner      uint16 = 0     // 新手
	Noblesse      uint16 = 1000  // 贵族
	Legend        uint16 = 2000  // 战神
	Even          uint16 = 2001  // 龙神
	Mercedes      uint16 = 2002  // 双弩精灵
	Phantom       uint16 = 2003  // 幻影
	Luminous      uint16 = 2004  // 夜光法师
	Shade         uint16 = 2005  // 隐月
	Citizen       uint16 = 3000  // 反抗者
	Demon         uint16 = 3001  // 恶魔
	Xenon         uint16 = 3002  // 尖兵
	WildHunter    uint16 = 3300  // 豹弩游侠
	Hayato        uint16 = 4001  // 剑豪
	Kanna         uint16 = 4002  // 阴阳师
	Mihile        uint16 = 5000  // 米哈尔
	Kaiser        uint16 = 6000  // 狂龙战士
	Zero          uint16 = 10000 // 神之子
	BeastTamer    uint16 = 11000 // 林之灵
	PinkBean      uint16 = 13000 // 品克缤
	Kinesis       uint16 = 14000 // 超能力者
	DragonWarrior uint16 = 17000 // 龙的传人
	Manager       uint16 = 800   // 管理员1
	GM            uint16 = 900   // 管理员2
	SuperGM       uint16 = 910   // 管理员3
)

func IsNotExtendSPJob(id uint16) bool {
	return id == BeastTamer || id == PinkBean || id == Manager || id == GM || id == SuperGM
}

func IsAdventurerJob(id uint16) bool {
	return id < 600
}

func IsDragonWarriorJob(id uint16) bool {
	return id == 17000 || id-17500 < 100
}

func IsCygnusJob(id uint16) bool {
	return id-1000 < 1000
}

func IsResistanceJob(id uint16) bool {
	return id-3000 < 1000
}

func IsEvanJob(id uint16) bool {
	return id-2200 < 100 || id == 2001
}

func IsMercedesJob(id uint16) bool {
	return id-2300 < 100 || id == 2002
}

func IsPhantomJob(id uint16) bool {
	return id-2400 < 100 || id == 2003
}

func IsLeaderJob(id uint16) bool {
	return id-5000 < 1000
}

func IsLuminousJob(id uint16) bool {
	return id-2700 < 100 || id == 2004
}

func IsDragonbornJob(id uint16) bool {
	// Evan
	return id-6000 < 1000
}

// is_wildhunter_job
func IsWildHunter(id uint16) bool {
	return id/100 == 33
}

func IsZeroJob(id uint16) bool {
	// return a1 == 10000 || a1 == 10100 || a1 == 10110 || a1 == 10111 || a1 == 10112;
	return id == 10000 || uint16(id-10100) < 100
}

func IsHiddenJob(id uint16) bool {
	// EunWol
	return id-2500 < 100 || id == 2005
}

func IsAranJob(id uint16) bool {
	return id-2100 < 100 || id == 2000
}

func IsKinesisJob(id uint16) bool {
	return id == 14000 || id == 14200 || id == 14210 || id == 14211 || id == 14212
}

func IsHayatoAndKannaJob(id uint16) bool {
	return id-4000 < 1000
}

func IsDamonJob(id uint16) bool {
	return id/100 == 31 || id == Demon
}

func IsXenonJob(id uint16) bool {
	return id/100 == 36 || id == Xenon
}

func IsBeastTamer(id uint16) bool {
	return id/100 == 112 || id == BeastTamer
}

func IsDualJob(id uint16) bool {
	return id/10 == 43
}

func IsKaiserJob(id uint16) bool {
	return id/100 == 61 || id == 6000
}

// is_beginner_job
func IsBeginnerJob(id uint16) bool {
	switch id {
	case 0,
		1000,
		2000,
		2001,
		2002,
		2003,
		2004,
		2005,
		3000,
		3001,
		3002,
		4001,
		4002,
		5000,
		6000,
		6001,
		8000,
		9000,
		10000,
		11000,
		13000,
		14000:
		return true
	}
	return false
}

// get_job_level
func GetJobLevel(id uint16) uint8 {
	var v2 uint16
	if IsBeginnerJob(id) || id%100 == 0 || id == 501 || id == 3101 {
		return 1
	}
	if IsEvanJob(id) {
		return GetEvanJobLevel(id)
	}
	if IsDualJob(id) {
		v2 = id % 10 / 2
	} else {
		v2 = id % 10
	}
	if v2 <= 2 {
		return uint8(v2 + 2)
	}
	return 0
}

func GetEvanJobLevel(id uint16) uint8 {
	var result uint8
	switch id {
	case 2200,
		2210:
		result = 1
	case 2211,
		2212,
		2213:
		result = 2
	case 2214,
		2215,
		2216:
		result = 3
	case 2217,
		2218:
		result = 4
	default:
		result = 0
	}
	return result
}
