package maple

type QuestState int8

const (
	QuestNotStart QuestState = iota
	QuestStarted
	QuestCompleted
)
