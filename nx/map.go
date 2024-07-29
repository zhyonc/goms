package nx

import "strconv"

type MapNX struct {
	Info    Info
	Lifes   []Life
	Portals []Portal
}

func (m *MapNX) GetWidth() int32 {
	return m.Info.VRRight - m.Info.VRLeft
}

func (m *MapNX) GetHeight() int32 {
	return m.Info.VRTop - m.Info.VRBottom
}

func (m *MapNX) GetLifeNPCID() uint32 {
	for _, l := range m.Lifes {
		if l.Type != "n" {
			continue
		}
		return l.GetID()
	}
	return 0
}

type Info struct {
	Bgm        string // count=87
	Cloud      int32  // count=87
	FieldLimit int32  // count=87
	// FieldScript       any    // count=73
	FieldType int32 // count=15
	// Fly               any   // count=87
	ForcedReturn int32 // count=87
	HideMinimap  bool  // count=87
	// LBBottom          any   // count=17
	// LBSide            any   // count=10
	// LBTop             any   // count=17
	// LimitSpeedAndJump any   // count=1
	// Link             any     // count=16
	MapDesc          string  // count=26
	MapMark          string  // count=87
	MobRate          float64 // count=87
	MoveLimit        int32   // count=48
	NoChair          bool    // count=28
	NoMapCmd         bool    // count=87
	OnFirstUserEnter bool    // count=87
	OnUserEnter      bool    // count=87
	// PartyStandAlone  any     // count=55
	ReturnMap int32 // count=87
	// StandAlone any   // count=55
	Swim      bool  // count=87
	TimeLimit int64 // count=1
	Town      bool  // count=87
	VRBottom  int32 // count=84
	VRLeft    int32 // count=84
	VRRight   int32 // count=84
	VRTop     int32 // count=84
	Version   int32 // count=87
}

type Life struct {
	Cy   int64 // count=53687
	F    bool  // count=53665 Face?
	Fh   int16 // count=53687 Foothold
	Hide bool  // count=51876
	// Hold        any   // count=77
	Id string // count=53687
	// Limitedname any   // count=301
	MobTime int64 // count=53620
	// Nofoothold  any   // count=43
	Rx0 int16 // count=53687
	Rx1 int16 // count=53687
	// Spine any   // count=99
	Type string // count=53686
	// UseDay      any   // count=393
	// UseNight    any   // count=393
	X int16 // count=53687
	Y int16 // count=53687
}

func (l *Life) GetID() uint32 {
	id, _ := strconv.Atoi(l.Id)
	return uint32(id)
}

func (l *Life) IsNPC() bool {
	return l.Type == "n"
}

func (l *Life) IsMob() bool {
	return l.Type == "m"
}

type Portal struct {
	// Delay            any    // count=80
	// HideTooltip      any    // count=80
	// HorizontalImpact any    // count=127
	// Image            any    // count=13
	// OnlyOnce         any    // count=80
	Pn string // count=280
	Pt int64  // count=280
	// ReactorName     any    // count=2
	Script string // count=87
	// SessionValue    any    // count=2
	// SessionValueKey any    // count=2
	// ShownAtMinimap  any    // count=15
	Tm int64  // count=280
	Tn string // count=280
	// VerticalImpact any    // count=2
	X int16 // count=280
	Y int16 // count=280
}
