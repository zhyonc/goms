package tip

const (
	WorldServerBusy Tip = 34  // 服務器用戶太多，連接出現延遲。請重新嘗試。
	WorldServerLag  Tip = 101 // 選擇的分流目前玩家過多而連綫上有發生延遲狀況。請選擇其他分流再進行游戲。
	MaxCapacity     Tip = 103 // 因伺服器的容納人數超過而無法使用此伺服器。請使用其他伺服器
)
