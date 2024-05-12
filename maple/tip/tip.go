package tip

type Tip int16

const (
	Success    Tip = 0
	DBFail     Tip = 6  // 因系統錯誤無法連缐！煩請稍後重新再試
	Unknown    Tip = 9  // 因系統錯誤無法連缐！煩請稍後重新再試
	ServerBusy Tip = 10 // 伺服器目前為忙碌狀態！請稍後再試
)

// const (
// 	Success Tip = iota
// 	TempBanned
// 	Banned                                // 請確認您的Gash帳號與游戲帳號是否為正常狀態
// 	Abandoned                             // 請確認您的Gash帳號與游戲帳號是否為正常狀態
// 	IncorrectPassword                     // 密碼輸入錯誤
// 	NotRegistered                         // 無此賬號！請確認後重新輸入
// 	DBFail                                // 因系統錯誤無法連缐！煩請稍後重新再試
// 	AlreadyConnected                      // 帳號使用中，目前是連缐！請重新確認
// 	NotConnectableWorld                   // 因系統錯誤無法連缐！煩請稍後重新再試
// 	Unknown                               // 因系統錯誤無法連缐！煩請稍後重新再試
// 	Timeout                               // 伺服器目前為忙碌狀態！請稍後再試
// 	NotAdult                              // 只限20歲以上的玩家登入
// 	AuthFail                              // Error 38
// 	ImpossibleIP                          // 您目前的IP無法以MASTER狀態來登入
// 	NotAuthorizedNexonID                  // 該賬號因“臨時加入”的期限已過而無法使用。若您未滿14歲，請登入gamania.com后，取得父母的同意，即可進行游戲
// 	NoNexonID                             // 這個Gash帳號並不存在
// 	IncorrectSSN2                         // Nothing happened
// 	WebAuthNeeded                         // 認證資料不符，請關閉目前開啓中的楓之谷官方網頁，並請重新再試
// 	DeleteCharacterFailedOnGuildMaster    // Nothing happened
// 	TempBlockedIP                         // 目前使用暫時被阻擋的IP連缐，詳細内容請洽官網的FAQ或是來電客服詢問
// 	IncorrectSPW                          // Nothing happened
// 	DeleteCharacterFailedEngaged          // Nothing happened
// 	SamePasswordAndSPW                    // Nothing happened
// 	WaitOTP                               // Nothing happened
// 	WrongOTP                              // Korean tip
// 	OverCountErrOTP                       // Korean tip
// 	SystemErr                             // Korean tip
// 	CancelInputDeleteCharacterOTP         // Nothing happened
// 	PaymentWarning                        // Korean tip
// 	DeleteCharacterFailedOnFamily         // Nothing happened
// 	InvalidCharacterName                  // Nothing happened
// 	IncorrectSSN                          // Nothing happened
// 	SSNConfirmFailed                      // Nothing happened
// 	SSNNotConfirmed                       // Nothing happened
// 	WorldTooBusy                          // 服務器用戶太多，連接出現延遲，請重新嘗試
// 	OTPReissuing                          // Korean tip
// 	OTPInfoNotExist                       // Korean tip
// 	Shutdowned                            // Korean tip
// 	DeleteCharacterFailedHasEntrustedShop // Nothing happened
// 	AlbaPerform                           // Nothing happened
// 	TransferredToNxEmailID                // Korean tip
// 	UntransferredToNxEmailID              // Korean tip
// 	RequestedMapleIDAlreadyInUse          // Nothing happened
// 	WaitSelectAccount                     // 登錄於Nexon信箱ID的楓之谷ID,選擇要連缐的楓之谷ID按下確認按鈕
// 	DeleteCharacterFailedProtectedItem    // Nothing happened
// 	UnauthorizedUser                      // Korean tip
// 	CannotCreateMoreMapleAccount          // Nothing happened
// 	CreateBanned                          // Nothing happened
// 	CreateTemporarilyBanned               // Nothing happened
// 	EventNewCharacterExpireFail           // Nothing happened
// 	SelectiveShotdowned                   // 根據青少年保護法未滿18歲的青少年在指定的選擇shot down時間限制使用所有缐上游戲
// 	NonownerRequest                       // Nothing happened
// 	OTPRequired                           // Korean tip
// 	GuestServiceClosed                    // Korean tip
// 	BlockedNexonID                        // Nothing happened
// 	DupMachineID                          // Nothing happened
// 	NotActiveAccount                      // Nothing happened
// 	IncorrectSPW4th                       // Nothing happened
// 	IncorrectSPW5th                       // Nothing happened
// 	InsufficientSPW                       // Nothing happened
// 	SameCharSPW                           // Nothing happened
// 	WebLaunchingOTPRequired               // Nothing happened
// 	MergeWorld_CreateCharacterBanned      // Nothing happened
// 	ChangeNewOTP                          // 因個人資訊保護安全原因將舊的U-OTP服務結束，以後無法在使用舊U-OTP來做認證
// 	BlockedByServiceArea                  // 針對海外網段登入進行阻擋，請經由(www.nexon.com)登入，確認本人資訊進行安全登入
// 	ExceedReservedDeleteCharacter         // Nothing happened
// 	UnionFieldChannelClosed               // Nothing happened
// 	UNK                                   // Nothing happened
// 	ProtectAccount                        // 保護模式設定時，無法使用楓之谷，因非正常的連狀態下而保護您的帳號，對其進行登入的限制：輸入多次密碼登入失敗，長時間沒有游戲連缐的狀態登入
// 	AntiMacroReq                          // Nothing happened
// 	AntiMacroCreateFailed                 // Nothing happened
// 	AntiMacroIncorrect                    // Nothing happened
// 	LimitCreateCharacter                  // Nothing happened
// 	ProtectSSOLogin                       // 重複輸入錯誤密碼 暫時無法登入 請利用尋找密碼來確認身份后再重新登入
// 	InvalidMapleIDThroughMobile           // Nothing happened
// 	InvalidPasswordThroughMobile          // Nothing happened
// 	HashedPasswordIsEmpty                 // Nothing happened
// 	NGS_For_Ass                           // Nothing happened
// 	AlreadyConnectedThroughMobile         // 帳號使用中
// 	Protected_For_Ass                     // Nothing happened
// 	Blocked_For_Ass                       // Nothing happened
// 	WrongVer                              // Nothing happened
// 	EMailVerify                           // 爲了完成加入會員，需要E-Mail的認證
// 	DenyCharacter                         // 此角色只能游玩到2017年11月28日為止
// 	InvalidObject                         // Nothing happened
// 	IncorrectLoginType_OtherToMapleID     // Nothing happened
// 	FailedUserCreate                      // Nothing happened
// 	MobileTokenInvalid                    // Nothing happened
// 	MobileTokenDeviceIDInvalid            // Nothing happened
// 	MobileTokenExpired                    // Nothing happened
// 	NotHaveNaverID                        // Nothing happened
// 	UserTossAIPlayer                      // Nothing happened
// 	InactivateMember                      // Nothing happened
// )
