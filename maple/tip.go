package maple

type Tip int16

const (
	TipSuccess                         Tip = 0
	TipTempBanned                      Tip = 1
	TipBanned                          Tip = 2 // 临时：您的帐号因为违反用户协议被封停到xxxx年xx月xx号 永久：由于您当前帐号违反法律法规或用户协议故停止服务。如有异议请联系客服热线：95105222
	TipAbandoned                       Tip = 3
	TipIncorrectPassword               Tip = 4 // 输入的密码有误，如果忘记密码，请点击找回ID/PW,重新设定密码后，再进行登录
	TipUsernameNotFound                Tip = 5 // 未登录的帐号，请重新确认
	TipDBFail                          Tip = 6 // 因系統錯誤無法連缐！煩請稍後重新再試
	TipAlreadyConnected                Tip = 7
	TipNotConnectableWorld             Tip = 8
	TipUnknown                         Tip = 9  // 因系統錯誤無法連缐！煩請稍後重新再試
	TipServerBusy                      Tip = 10 // 目前因链接邀请过多 服务器未能处理 请稍后再尝试
	TipNotAdult                        Tip = 11
	TipImpossibleIP                    Tip = 13
	TipTempBlockedIP                   Tip = 19
	TipSamePasswordAndSPW              Tip = 22
	TipWorldTooBusy                    Tip = 34
	TipMergeWorldCreateCharacterBanned Tip = 62
	TipBlockedByServiceArea            Tip = 64
	TipProtectSSOLogin                 Tip = 73
	TipEMailVerify                     Tip = 82
)
