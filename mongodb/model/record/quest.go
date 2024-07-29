package record

import "goms/maple"

type Quest struct {
	ID    int32            `bson:"id"`
	State maple.QuestState `bson:"state"`
}
