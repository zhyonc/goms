package opcode

const (
	CClientSocket_SendSecurityPacket                 = 102
	CClientSocket_SendPermissionRequest              = 103 // Done
	CLogin_CheckLoginAuthInfo                        = 105 // Done
	CLogin_SendSelectWorldRequest                    = 106 // Done
	CClientSocket_MigrateIn                          = 110
	CLogin_SendSelectCharacterRequest                = 111 // Done
	CLogin_ChangeStepImmediate                       = 113
	CLogin_SendCheckDuplicateIDPacket                = 116 // Done
	CLogin_GotoWorldSelect                           = 117 // Done
	CLogin_SendNewCharPacket                         = 125 // Done
	CLogin_SendDeleteCharPacket                      = 128 // Done
	CLogin_SendReservedDeleteCharacterConfirmStep    = 129 // Done
	CLogin_SendReservedDeleteCharacterCancelStep     = 130 // Done
	CClientSocket_SendExceptionLog                   = 133 // Done
	CClientSocket_OnAliveReq_Callback                = 148 // Todo
	CClientSocket_OnCheckAliveAck_Callback           = 149 // Todo
	CWvsApp_SendBackupPacket                         = 150 // Done Client dump log
	CClientSocket_CRCErrorLog                        = 151
	CClientSocket_ApplyHotfix                        = 153 // Done
	CTerminateException                              = 154 // UNK
	CrashReporter_PreCallback                        = 156 // Done
	CLogin_ProgressNewCharStep                       = 157 // UNK
	CLogin_DirectGoToField                           = 159 // Done
	CLogin_SendChangeCharOrderRequest                = 160 // Done
	CLogin_SendGenderSetRequest                      = 163 // Done
	CLogin_SendServerStatusRequest                   = 171
	CLogin_SendNewCharPacket_ClientToGame            = 178
	CLogin_SendWorldStatusCheck                      = 180 // Done
	CLogin_ApplyRSAKey                               = 181 // Done
	CLogin_OnCreateCharStep_Callback                 = 184 // Done
	CUserLocal_TalkToNpc                             = 235 // ACK
	CUserLocal_TalkToNpcStep                         = 237 // ACK Found in CScriptMan::OnScriptMessage
	CWvsContext_SendStatChangeItemUseRequest         = 260 // ACK
	CWvsContext_SendMobSummonItemUseRequest          = 263 // ACK
	CWvsContext_SendScriptRunItemRequest             = 266 // ACK
	CWvsContext_SendRecipeOpenItemRequest            = 267 // ACK
	CWvsContext_SendShopScannerItemUseRequest        = 285 // ACK
	UNK340                                           = 340
	CUserLocal_SendClientResolution                  = 358 // Done
	CWvsContext_RequestInstanceTable                 = 364 // Done
	CWvsContext_OnCheckWeddingExResult_Callback      = 403 // ACK
	CWvsContext_OnPartyRequest_Callback              = 405 // ACK
	CWvsContext_OnPartyRequest_Callback2             = 408 // LABEL_247
	CFuncKeyMappedMan_Handler429                     = 429 // ACK ChangePetConsumeItemID/ChangePetConsumeMPItemID/SaveFuncKeyMap
	CWvsContext_OnMarriageRequest_Callback           = 434
	CWvsContext_OnFamilySummonRequest_Callback       = 449  // ACK
	CMobPool_OnMobCrcKeyChanged_Callback             = 457  // ACK
	CMob_OnCrcDataRequest_Callback                   = 458  // ACK
	CUserLocal_CheckNpcSpeechQuest                   = 523  // ACK
	CWvsContext_OnUserDamageOnFallingCheck_Callback  = 528  // ACK
	CWvsContext_OnPersonalShopBuyCheck_Callback      = 529  // ACK
	CWvsContext_SendStatChangeItemUseRequestByPetQ   = 566  // ACK
	CQuickslotKeyMappedMan_SaveQuickslotKeyMap       = 600  // ACK CQuickslotKeyMappedMan::OnInit->SaveQuickslotKeyMap
	CClientSocket_OnPingCheckResult_Callback         = 620  // ACK
	CWvsContext_OnMonsterBattleSystemResult_Callback = 631  // ACK
	CLogin_UNK707                                    = 707  // UNK
	CWvsContext_OnInventoryOperation_Callback        = 823  // ACK
	CUIRecipeSummon_OnButtonClicked                  = 939  // ACK
	CWvsContext_UI_Open_Callback                     = 1152 // case 597
)
