package maple

const ClassCount uint8 = 22 // Current version all job count

const (
	ClassResistance    uint32 = iota // 反抗者
	ClassAdventures                  // 冒险家
	ClassKnightCygnus                // 冒险骑士团
	ClassAran                        // 战神
	ClassEvan                        // 龙神
	ClassMercedes                    // 双弩精灵
	ClassDemon                       // 恶魔
	ClassPhantom                     // 幻影
	ClassDualBlade                   // 暗影双刀
	ClassMikhail                     // 米哈尔
	ClassLuminous                    // 夜光法师
	ClassKaiser                      // 狂龙战士
	ClassAngelicBuster               // 爆莉萌天使
	ClassCannonShooter               // 火炮手
	ClassXenon                       // 尖兵
	ClassZero                        // 神之子
	ClassEunWol                      // 隐月
	ClassPinkBean                    // 品克缤
	ClassKinesis                     // 超能力者
	ClassHayato        uint32 = 1001 // 剑豪
	ClassKanna         uint32 = 1002 // 阴阳师
	ClassBeastTamer    uint32 = 1003 // 林之灵
	ClassDragonWarrior uint32 = 1004 // 龙的传人-2021年2月4日->墨玄
)

var EnableClassCreation map[uint32]struct{} = map[uint32]struct{}{
	ClassAdventures: {},
}
