package character

type TamingMob struct {
	Level   uint32 `bson:"level"`
	Exp     uint32 `bson:"exp"`
	Fatigue uint32 `bson:"fatigue"`
}
