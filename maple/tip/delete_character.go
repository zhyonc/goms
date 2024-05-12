package tip

const (
	UnknownReson            Tip = 9   // 因不明原因動作失效！請稍後再試
	IncorrectSecondPassword Tip = 16  // 第2組密碼錯誤 other value=20
	IsGuildMaster           Tip = 18  // 公會長身份無法刪除角色，請解散公會後再重試
	IsGetMarried            Tip = 21  // 無法刪除訂婚及訂婚程序進行中的角色
	IsGuildMember           Tip = 29  // 無法刪除有家族的角色，請先退出家族後重新再試
	HaveShopItem            Tip = 38  // 因交易的道具尚未取回而無法刪除角色。請跟自由市場NPC富蘭德領回所有物品
	HoldSealLock            Tip = 44  // 擁有使用封印之鎖的角色無法刪除
	IncorrectCount4         Tip = 57  // 第二組密碼輸入錯誤達4次。若輸入錯誤的密碼5次，將會設定爲保護模式，部分服務使用將受限
	IncorrectCount5         Tip = 58  // BackToLoginScreen 第二組密碼輸入錯誤達5次, 已設定成保護模式，部分服務使用將受限。
	HavaAuctionItem         Tip = 93  // 在楓之谷有登記拍賣物品的角色無法刪除。請結束拍賣交易後再進行刪除。
	InAllianceBattleField   Tip = 96  // 無法刪除包含在楓之谷聯盟戰地的角色
	InMigration             Tip = 122 // 伺服器移民處理中，無法刪除角色
	HoldSealItem            Tip = 124 // 角色身上有被封印的道具！目前無法刪除該角色
)
