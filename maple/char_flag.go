package maple

type CharFlag uint64

const (
	FlagNone               CharFlag = 0
	FlagCharacter          CharFlag = 0x1
	FlagMoney              CharFlag = 0x2
	FlagItemSlot           CharFlag = 0x2
	FlagItemSlotEquip      CharFlag = 0x4
	FlagReturnEffectInfo   CharFlag = 0x4
	FlagItemSlotConsume    CharFlag = 0x8
	FlagDressUpInfo        CharFlag = 0x8
	FlagItemSlotSetUp      CharFlag = 0x10
	FlagEvolutionInfo      CharFlag = 0x10
	FlagItemSlotEtc        CharFlag = 0x20
	FlagItemSlotCash       CharFlag = 0x40
	FlagInventorySize      CharFlag = 0x80
	FloagMemorialCubeInfo  CharFlag = 0x80
	FlagSkillRecord        CharFlag = 0x100
	FlagQuestRecord        CharFlag = 0x200
	FlagMiniGameRecord     CharFlag = 0x400
	FlagLikePoint          CharFlag = 0x400
	FlagRingRecord         CharFlag = 0x800
	FlagZeroInfo           CharFlag = 0x800
	FlagMapTransfer        CharFlag = 0x1000
	FlagAvatar             CharFlag = 0x2000
	FlagQuestComplete      CharFlag = 0x4000
	FlagSkillCooltime      CharFlag = 0x8000
	FlagMonsterBattleInfo  CharFlag = 0x8000
	FlagRunnerGameRecord   CharFlag = 0x20000
	FlagQuestRecordEx      CharFlag = 0x40000
	FlagMonsterCollection  CharFlag = 0x40000
	FlagFamiliar           CharFlag = 0x80000
	FlagPendantExt         CharFlag = 0x100000
	FlagSoulCollection     CharFlag = 0x100000
	FlagWildHunterInfo     CharFlag = 0x200000
	FlagRedLeafInfo        CharFlag = 0x200000
	FlagFarmPotential      CharFlag = 0x200000
	FlagCoreAura           CharFlag = 0x400000
	FlagItemPot            CharFlag = 0x800000
	FlagCoreInfo           CharFlag = 0x1000000
	FlagExpConsumeItem     CharFlag = 0x2000000
	FlagPotionPot          CharFlag = 0x4000000
	FlagShopBuyLimit       CharFlag = 0x4000000
	FlagChosenSkills       CharFlag = 0x10000000
	FlagStolenSkills       CharFlag = 0x20000000
	FlagDayLimit           CharFlag = 0x40000000
	FlagCharacterPotential CharFlag = 0x80000000
	FlagAll                CharFlag = 0xFFFFFFFFFFFFFFFF
)

type ZeroFlag uint16

const (
	ZeroFlagBeta                   ZeroFlag = 0x1
	ZeroFlagSubHP                  ZeroFlag = 0x2
	ZeroFlagSubMP                  ZeroFlag = 0x4
	ZeroFlagSubSkin                ZeroFlag = 0x8
	ZeroFlagSubHair                ZeroFlag = 0x10
	ZeroFlagSubFace                ZeroFlag = 0x20
	ZeroFlagSubMHP                 ZeroFlag = 0x40
	ZeroFlagSubMMP                 ZeroFlag = 0x80
	ZeroFlagDBCharZeroLinkCashPart ZeroFlag = 0x100
	ZeroFlagHairColor              ZeroFlag = 0x200
	ZeroFlagAll                    ZeroFlag = 0xFFFF
)
