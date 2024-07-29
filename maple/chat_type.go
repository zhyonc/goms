package maple

type ChatType int16

const (
	CT_Normal ChatType = iota
	CT_Whisper
	CT_GroupParty
	CT_GroupFriend
	CT_GroupGuild
	CT_GroupAlliance
	CT_GameDesc
	CT_Tip
	CT_Notice
	CT_Notice2
	CT_AdminChat
	CT_SystemNotice
	CT_SpeakerChannel
	CT_SpeakerWorld
	CT_SpeakerWorldGuildSkill
	CT_ItemSpeaker
	CT_ItemSpeakerItem
	CT_SpeakerBridge
	CT_SpeakerWorldExPreview
	CT_Mob
	CT_Expedition
	CT_ItemMessage
	CT_MiracleTime
	CT_LotteryItemSpeaker
	CT_LotteryItemSpeakerWorld
	CT_AvatarMegaphone
	CT_PickupSpeakerWorld
	CT_WorldName
	CT_BossArenaNotice
	CT_Claim
	CT_AfreecaTv
	// non kmst from here
	CT_GachaReward
	CT_GachaRed
	CT_GachaRed2 // same as GachaRed(32)
	CT_DarkBlue2 // same as ItemSpeakerItem(16)
	CT_ItemNoItemSmegaDarkText
	CT_WhiteOnGreen
	CT_CakeSpeaker
	CT_PieSpeaker
	CT_BlackOnWhite
)
