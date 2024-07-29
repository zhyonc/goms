package maple

type ExpedtionRetCode int8

const (
	ExpNoti_Get1           ExpedtionRetCode = '\\'
	ExpNoti_Get2           ExpedtionRetCode = '^'
	ExpNoti_Get3           ExpedtionRetCode = '`'
	ExpNoti_Removed1       ExpedtionRetCode = ']'
	ExpNoti_Removed2       ExpedtionRetCode = 'd'
	ExpNoti_Removed3       ExpedtionRetCode = 'f'
	ExpNoti_Removed4       ExpedtionRetCode = 'g'
	ExpNoti_Notice1        ExpedtionRetCode = '_'
	ExpNoti_Notice2        ExpedtionRetCode = 'c'
	ExpNoti_Notice3        ExpedtionRetCode = 'e'
	ExpNoti_UNK            ExpedtionRetCode = 'a'
	ExpNoti_MasterChanged  ExpedtionRetCode = 'h'
	ExpNoti_Modified       ExpedtionRetCode = 'i'
	ExpNoti_Invite         ExpedtionRetCode = 'k'
	ExpNoti_ResponseInvite ExpedtionRetCode = 'l'
)
