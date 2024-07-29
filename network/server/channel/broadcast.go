package channel

import (
	"goms/maple"
	"goms/mongodb/model/inventory"
)

type Broadcast struct {
	Type        maple.BroadcastType
	Msg         string
	NextMsg     []string
	Msg2        string
	Msg3        string
	WhisperIcon bool
	Item        inventory.ItemSlotBundle
	Timeout     uint32
	NPCID       uint32
	QuestID     uint32
	CharID      uint32
	Width       uint32
	Height      uint32
}

func NewBroadcast(msg string) Broadcast {
	b := Broadcast{
		Type: maple.SlideNotice,
		Msg:  msg,
	}
	return b
}

func (b *Broadcast) SetMegaphone(isWhisper bool) {
	b.Type = maple.Megaphone
	b.WhisperIcon = isWhisper
}

func (b *Broadcast) SetItemMegaphone(isWhisper bool, item inventory.ItemSlotBundle) {
	b.Type = maple.ItemMegaphone
	b.WhisperIcon = isWhisper
	b.Item = item
}

func (b *Broadcast) SetTripleMegaphone(isWhisper bool, msgs ...string) {
	b.Type = maple.TripleMegaphone
	b.WhisperIcon = isWhisper
	for _, msg := range msgs {
		if msg == "" {
			continue
		}
		b.NextMsg = append(b.NextMsg, msg)
	}
}

func (b *Broadcast) SetBlowWeather(item inventory.ItemSlotBundle) {
	b.Type = maple.BlowWeather
	b.Item = item
}

func (b *Broadcast) SetBalloonMessage(item inventory.ItemSlotBundle, timeout uint32) {
	b.Type = maple.BalloonMessage
	b.Item = item
	b.Timeout = timeout
}

func (b *Broadcast) SetWhiteYellowItemInfo(item inventory.ItemSlotBundle) {
	b.Type = maple.WhiteYellow_ItemInfo
	b.Item = item
}

func (b *Broadcast) SetYellow(item inventory.ItemSlotBundle) {
	b.Type = maple.Yellow
	b.Item = item
}

func (b *Broadcast) SetBlueChatItemInfo(item inventory.ItemSlotBundle) {
	b.Type = maple.BlueChatItemInfo
	b.Item = item
}

func (b *Broadcast) SetGMErrorMessage(npcID uint32) {
	b.Type = maple.GMErrorMessage
	b.NPCID = npcID
}

func (b *Broadcast) SetYellowChatFiledItemInfo(item inventory.ItemSlotBundle) {
	b.Type = maple.YellowChatFiledItemInfo
	b.Item = item
}

func (b *Broadcast) SetTryRegisterAutoStartQuest(questID uint32, timeout uint32) {
	b.Type = maple.TryRegisterAutoStartQuest
	b.QuestID = questID
	b.Timeout = timeout
}

func (b *Broadcast) SetTryRegisterAutoStartQuestNoAnnouncement(questID uint32) {
	b.Type = maple.TryRegisterAutoStartQuest
	b.QuestID = questID
}

func (b *Broadcast) SetRedWithChannelInfo(charID uint32) {
	b.Type = maple.RedWithChannelInfo
	b.CharID = charID
}

func (b *Broadcast) SetPopUpNotice(width, height uint32) {
	b.Type = maple.PopUpNotice
	b.Width = width
	b.Height = height
}
