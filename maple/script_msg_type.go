package maple

type ScriptMsgType int8

const (
	Say ScriptMsgType = iota
	SayUNK
	SayImage
	AskYesNo
	AskText
	AskNumber
	AskMenu
	InitialQuiz
	InitialSpeedQuiz
	ICQuiz
	AskAvatar
	AskAndroid
	AskPet
	AskPetAl
	AskActionPetEvolution
	_
	AskYesNo2
	_
	AskBoxText
	AskSlideMenu
	_
	_
	_
	_
	AskSelectMenu
	AskAngelicBuster
	SayIllustration
	SayIllustration2
	AskYesNoIllustration
	AskYesNoIllustration2
	AskMenuIllustration
	AskYesNoIllustration3
	AskYesNoIllustration4
	AskMenuIllustration2
	AskAvatarZero
	_
	AskWeaponBox
	AskBoxText_BgImg
	AskUserSurvey
	_
	AskMixHair
	AskMixHairExZero
	AskCustomMixHair
	AskCustomMixHairAndProb
	AskMixHairNew
	AskMixHairNewExZero
	_
	AskScreenShinningStarMsg
	_
	_
	AskNumberUseKeyPad
	SpinOffGuitarRhythmGame
	GhostParkEnter
)
