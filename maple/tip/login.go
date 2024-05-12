package tip

const (
	TempBanned                      Tip = 1  // xxxx年x月x日xx點xx分之後開始可以登入
	Banned                          Tip = 2  // 請確認您的Gash帳號與游戲帳號是否為正常狀態或寫下封禁原因
	Abandoned                       Tip = 3  // 請確認您的Gash帳號與游戲帳號是否為正常狀態
	IncorrectPassword               Tip = 4  // 密碼輸入錯誤
	NotRegistered                   Tip = 5  // 無此賬號！請確認後重新輸入
	AlreadyConnected                Tip = 7  // 帳號使用中，目前是連缐！請重新確認
	NotConnectableWorld             Tip = 8  // 因系統錯誤無法連缐！煩請稍後重新再試
	NotAdult                        Tip = 11 // 只限20歲以上的玩家登入
	ImpossibleIP                    Tip = 13 // 您目前的IP無法以MASTER狀態來登入
	TempBlockedIP                   Tip = 19 // 目前使用暫時被阻擋的IP連缐，詳細内容請洽官網的FAQ或是來電客服詢問
	SamePasswordAndSPW              Tip = 22
	WorldTooBusy                    Tip = 34 // 服務器用戶太多，連接出現延遲，請重新嘗試
	MergeWorldCreateCharacterBanned Tip = 62 // 在該伺服器限制使用整合賬號暫時只能創建1個角色.因已經持有角色所以無法再創建角色了
	BlockedByServiceArea            Tip = 64 // 針對海外網段登入進行阻擋，請經由(www.nexon.com)登入，確認本人資訊進行安全登入。Will exit game
	ProtectSSOLogin                 Tip = 73 // 重複輸入錯誤密碼 暫時無法登入 請利用尋找密碼來確認身份后再重新登入
	EMailVerify                     Tip = 82 // 爲了完成加入會員，需要E-Mail的認證
)
