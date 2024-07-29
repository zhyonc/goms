package maple

type MessageType int8

const (
	DropPickUpMessage MessageType = iota
	QuestRecordMessage
	QuestRecordMessageAddValidCheck
	CashItemExpireMessage
	IncEXPMessage
	IncSPMessage
	IncPOPMessage
	IncMoneyMessage
	IncGPMessage
	IncCommitmentMessage
	GiveBuffMessage
	GeneralItemExpireMessage
	SystemMessage
	QuestRecordExMessage
	WorldShareRecordMessage
	ItemProtectExpireMessage
	ItemExpireReplaceMessage
	ItemAbilityTimeLimitedExpireMessage
	SkillExpireMessage
	IncNonCombatStatEXPMessage
	LimitNonCombatStatEXPMessage
	_
	AndroidMachineHeartAlsetMessage
	IncFatigueByRestMessage
	IncPvPPointMessage
	PvPItemUseMessage
	WeddingPortalError
	IncHardCoreExpMessage
	NoticeAutoLineChanged
	EntryRecordMessage
	EvolvingSystemMessage
	EvolvingSystemMessageWithName
	CoreInvenOperationMessage
	NxRecordMessage
	BlockedBehaviorMessage
	IncWPMessage
	MaxWPMessage
	StylishKillMessage
	BarrierEffectIgnoreMessage
	ExpiredCashItemResultMessage
	CollectionRecordMessage
	RandomChanceMessage
	UNK42
	UNK43
	UNK44
)
