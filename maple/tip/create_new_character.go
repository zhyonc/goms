package tip

const (
	AlreadyConnectedAndLogout Tip = 7    // BackToLoginScreen 帳號使用中，目前是連缐！請重新確認
	NotEnoughSlot             Tip = 9    // 角色欄位不足，無法再創建
	InvalidCharacterName      Tip = 30   // 您無法使用該昵稱，請選擇其它名稱后再次輸入
	DenyJob                   Tip = 99   // 無法選擇此職業
	ServerLimitCreate         Tip = 105  // 限制創角的伺服器。請使用其他伺服器
	InBurningServer           Tip = 108  // 燃燒、燃燒2、燃燒3伺服器當中無法同時創建角色
	InYetPinkBeanServer       Tip = 119  // 在雪吉拉x皮卡啾伺服器無法同時創建角色
	Migranting                Tip = 122  // 伺服器移民申請中！該賬號無法新建角色
	CreateUnknown             Tip = -124 // 因未知錯誤導致失效！請稍後再試
)
