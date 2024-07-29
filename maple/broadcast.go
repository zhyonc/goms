package maple

type BroadcastType uint8

const (
	Notice BroadcastType = iota
	PopUpMessage
	DarkBlueOnLightBlue
	Megaphone
	SlideNotice
	PartyChat
	BlueChatItemInfo
	GMErrorMessage
	ItemMegaphone // Holds item info
	ItemMegaphoneNoItem
	TripleMegaphone
	YellowChatFiledItemInfo //  item shown when clicked,  does hold Item info
	BlowWeather
	TryRegisterAutoStartQuest                // tries to auto start quests with the announcement (?) - probably the  "A quest has arrived! Please clock on the icon at the botfom of your screen."
	TryRegisterAutoStartQuest_NoAnnouncement // tries to auto start quests (?)
	SwedishFlag                              // Repeats the string 3x on the same line
	RedWithChannelInfo                       // May be for  /find ?
	WhiteYellow_ItemInfo                     // Holds item info
	BlueChatItemInfo2
	WhiteYellow
	PopUpNotice
	Yellow // Holds item info
	Yellow2
	MegaphoneNoMessage
	BalloonMessage
)
