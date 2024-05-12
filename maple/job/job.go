package job

const JobOrder uint8 = 240 // UI.wz/Login.img/RaceSelect_new/order

type JobID uint16

const (
	Beginner      JobID = 0     // 新手
	Noblesse      JobID = 1000  // 貴族
	Legend        JobID = 2000  // 狂狼勇士
	Even          JobID = 2001  // 龍魔導士
	Mercedes      JobID = 2002  // 精靈遊俠
	Phantom       JobID = 2003  // 幻影俠盜
	Luminous      JobID = 2004  // 夜光
	Shade         JobID = 2005  // 隱月
	Citizen       JobID = 3000  // 末日反抗軍
	Demon         JobID = 3001  // 惡魔
	Xenon         JobID = 3002  // 傑諾
	Hayato        JobID = 4001  // 劍豪
	Kanna         JobID = 4002  // 陰陽師
	Mihile        JobID = 5000  // 米哈逸
	Kaiser        JobID = 6000  // 凱撒
	AngelicBuster JobID = 6001  // 天使破壞者
	Cadena        JobID = 6002  // 卡蒂娜
	Kain          JobID = 6003  // 凱殷
	Zero          JobID = 10000 // 神之子
	BeastTamer    JobID = 11000 // 幻獸師
	PinkBean      JobID = 13000 // 皮卡啾
	Yeti          JobID = 13001 // 雪吉拉
	Kinesis       JobID = 14000 // 凱內西斯
	Illium        JobID = 15000 // 伊利恩
	Ark           JobID = 15001 // 亞克
	Adele         JobID = 15002 // 阿戴爾
	HoYoung       JobID = 16000 // 虎影
	Lara          JobID = 16001 // 菈菈
	MoXuan        JobID = 17000 // 墨玄
	Manager       JobID = 800
	GM            JobID = 900
	SuperGM       JobID = 910
)
