package class

type ClassID uint16

const (
	Resistance    ClassID = iota // 末日反抗軍
	Adventures                   // 冒險家
	KnightCygnus                 // 皇家騎士團
	Aran                         // 狂狼勇士
	Evan                         // 龍魔導士
	Mercedes                     // 精靈游俠
	Demon                        // 惡魔
	Phantom                      // 幻影俠盜
	DualBlade                    // 影武者
	Mikhail                      // 米哈逸
	Luminous                     // 夜光
	Kaiser                       // 凱撒
	AngelicBuster                // 天使破壞者
	CannonShooter                // 重炮指揮官
	Xenon                        // 傑諾
	Zero                         // 神之子
	EunWol                       // 隱月
	PinkBean                     // 皮卡啾
	Kinesis                      // 凱内西斯
	Cadena                       // 卡蒂娜
	Illium                       // 伊利恩
	Ark                          // 亞克
	Pathfinder                   // 開拓者
	HoYoung                      // 虎影
	Adele                        // 阿戴爾
	Kain                         // 凱殷
	Yeti                         // 雪吉拉
	Lara                         // 菈菈
	MoXuan        = 1000         // 墨玄
	Hayato        = 1001         // 劍豪
	Kanna         = 1002         // 陰陽師
	BeastTamer    = 1003         // 幻獸師
)

var EnableClassCreation map[ClassID]struct{} = map[ClassID]struct{}{
	Lara:       {},
	Adventures: {},
}
