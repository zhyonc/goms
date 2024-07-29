package manager

type GlobalVar string

const (
	NPCManagerName   GlobalVar = "cm"
	QuestManagerName GlobalVar = "qm"
)

type Manager interface {
	Dispose()
}

type NPCManager interface {
	Manager
	SendNext(text string)
	SendNextPrev(text string)
	SendPrev(text string)
	Warp(mapID, portalID int32)
}

type QuestManager interface {
	Manager
}
