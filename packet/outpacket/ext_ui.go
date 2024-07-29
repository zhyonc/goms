package outpacket

import (
	"goms/maple"
	"goms/nx"
	"goms/util"
)

// Call by CWvsContext::OnRequestEventList
// CWvsContext::OnLiveEvent
func CWvsContextOnLiveEvent(p *outPacket) {
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		LiveEventEncode(p)
	}
}

// Call by CWvsContext::OnLiveEvent
// LiveEvent::LIVE_EVENT::Decode
func LiveEventEncode(p *outPacket) {
	p.EncodeStr("")  // sName
	p.EncodeInt32(0) // nCategory
	p.EncodeInt32(0) // nEventType
	p.EncodeInt32(0) // nEventValue
	p.EncodeInt32(0) // nDateStart
	p.EncodeInt32(0) // nDateEnd
	p.EncodeInt32(0) // nTimeStart
	p.EncodeInt32(0) // nTimeEnd
	p.EncodeInt32(0) // nMinLevel
	p.EncodeInt32(0)
	p.EncodeBool(false)
	p.EncodeInt32(0)
	p.EncodeInt32(0) // count?
	p.EncodeInt32(0) // count?
	p.EncodeStr("")  // sDesc
	p.EncodeStr("")
}

// Call by CWvsContext::OnRequestEventList
// CWvsContext::MakeEventListImg
func CWvsContextMakeEventListImg(p *outPacket) bool {
	p.EncodeStr("") // sMain
	condition := false
	p.EncodeBool(condition)
	if condition {
		p.EncodeInt32(0) // maxCount?
	}
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		p.EncodeByte(0) // p_pvarg?
		p.EncodeStr("") // nAlarmQuestIdx
	}
	count = 0
	p.EncodeUint32(uint32(count)) // nEntryCount
	for i := 0; i < count; i++ {
		p.EncodeInt32(0) // p_pvarg?
		p.EncodeStr("")  // sName
		p.EncodeStr("")  // sCollectTip
		p.EncodeInt32(0) // timeStart
		p.EncodeInt32(0) // timeEnd
		p.EncodeInt32(0) // dateStart
		p.EncodeInt32(0) // dateEnd
		p.EncodeInt32(0) // UI
		p.EncodeInt32(0) // prior
		p.EncodeByte(0)  // hot
		p.EncodeByte(0)  // expEvent
		p.EncodeByte(0)  // attend
		p.EncodeByte(0)  // invisible
		p.EncodeByte(0)  // continue
		count := 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeInt32(0) // reward
		}
		count = 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeInt32(0) // quest
		}
		condition = false
		p.EncodeBool(condition)
		if condition {
			count := 0
			p.EncodeUint32(uint32(count)) // nOpenUIIndex
			for i := 0; i < count; i++ {
				p.EncodeStr("") // sAlarmQuestRange
			}
		}
		p.EncodeStr("")  // sTooltip
		p.EncodeStr("")  // sResourcePath
		p.EncodeInt32(0) // nOpenUIIndex
		p.EncodeByte(0)  // sAlarmQuestIdx
		condition := false
		p.EncodeBool(condition)
		if condition {
			p.EncodeInt32(0) // lvmin
			p.EncodeInt32(0) // lvmax
			p.EncodeInt32(0) // world
		}
	}
	return false
}

// Call by CStage::OnSetField
// CWvsContext::LogoutEvent
func CWvsContextLogoutEvent(p *outPacket) {
	p.EncodeInt32(0) // UNK
	for i := 0; i < 3; i++ {
		p.EncodeInt32(0) // UNk
	}
}

// Call by CStage::OnSetField
// CWvsContext::DecodeStarPlanetRoundInfo
func CWvsContextEncodeStarPlanetRoundInfo(p *outPacket) {
	p.EncodeUint32(0)         // m_nStarPlanetRoundID
	p.EncodeByte(0)           // m_nStarPlanetRoundState
	p.EncodeFT(util.ZeroTime) // m_ftStarPlanetRoundEndDate
}

// Call by CWvsContext::OnMessage
// CWvsContext::OnQuestRecordMessage
func CWvsContextQuestRecordMessage(p *outPacket, quest *nx.QuestNX) {
	p.EncodeUint32(quest.ID)        // nQRKey questID
	p.EncodeByte(byte(quest.State)) // nState
	CWvsContextQuestFlashByQRMessage(p, quest.State)
}

// Call by CWvsContext::OnMessage
// CWvsContext::OnQuestRecordExMessage
func CWvsContextQuestRecordExMessage(p *outPacket, quest *nx.QuestNX) {
	p.EncodeUint32(quest.ID) // nQRKey questID
	p.EncodeStr("")          // sValue
}

// Call by CWvsContext::OnQuestRecordMessage
// CWvsContext::QuestFlashByQRMessage
func CWvsContextQuestFlashByQRMessage(p *outPacket, state maple.QuestState) {
	switch state {
	case maple.QuestNotStart:
		p.EncodeBool(false) //  If quest is completed, but should never be true?
	case maple.QuestStarted:
		p.EncodeStr("") // custom data
	case maple.QuestCompleted:
		p.EncodeFT(util.ZeroTime) // quest completed date
	}
}
