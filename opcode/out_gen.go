// Code generated by opcode_test, DO NOT EDIT.
package opcode

var OutMap map[uint16]string = map[uint16]string{
	0: "CLogin_OnCheckPasswordResult",
	1: "CLogin_OnWorldInformation",
	2: "CLogin_OnLatestConnectedWorld",
	3: "CLogin_OnRecommendWorldMessage",
	4: "CLogin_OnSetClientKey",
	5: "CLogin_OnSetPhysicalWorldID",
	6: "CLogin_OnSelectWorldResult",
	7: "CLogin_OnSelectCharacterResult",
	8: "CLogin_OnAccountInfoResult",
	9: "CLogin_OnCreateMapleAccountResult",
	10: "CLogin_OnCheckDuplicatedIDResult",
	11: "CLogin_OnCreateNewCharacterResult",
	12: "CLogin_OnDeleteCharacterResult",
	13: "CLogin_OnReservedDeleteCharacterResult",
	14: "CLogin_OnReservedDeleteCharacterCancelResult",
	15: "CLogin_OnRenameCharacterResult",
	24: "CLogin_OnChangeSPWResult",
	25: "CLogin_OnCheckSPWExistResult",
	26: "CLogin_OnCheckWebLoginEmailID",
	27: "CLogin_OnCheckDeleteCharacterOTP",
	29: "CLogin_OnAlbaResult",
	30: "_",
	32: "CLogin_OnSetAutoSelectedWorld",
	37: "CLogin_OnCheckSPWResult",
	38: "_",
	39: "CLogin_OnChooseGender",
	40: "CLogin_OnGenderSetResult",
	41: "_",
	42: "_",
	43: "_",
	44: "_",
	45: "_",
	46: "_",
	47: "_",
	54: "CLogin_OnWorldStatus",
	55: "CLogin_OnRSAKey",
	56: "_",
	57: "_",
	58: "CLogin_OnCreateCharStep",
	59: "_",
	1485: "CLogin_OnCheckSPWOnCreateNewCharacterResult",
	1486: "CLogin_OnSetSPWResul",
	17: "CClientSocket_OnMigrateCommand",
	18: "CClientSocket_OnAliveReq",
	19: "CClientSocket_OnChatServerResult",
	20: "CClientSocket_OnPingCheckResult",
	21: "CClientSocket_OnAuthenCodeChanged",
	22: "CClientSocket_OnAuthenMessage",
	23: "CClientSocket_OnSecurityPacket",
	31: "CClientSocket_OnCheckAliveAck",
	34: "CClientSocket_OnReceiveHotfix",
	35: "CClientSocket_UNK35",
	48: "CClientSocket_UNK48",
	60: "CMapLoadable_OnSetBackEffect",
	61: "CMapLoadable_OnSetMapTagedObjectVisible",
	62: "CMapLoadable_OnSetMapTaggedObjectSmoothVisible",
	63: "CMapLoadable_OnSetMapTaggedObjectAnimation",
	64: "CMapLoadable_OnSetMapObjectAnimation",
	65: "CMapLoadable_OnSetMapObjectVisible",
	66: "CMapLoadable_OnSetMapObjectMove",
	67: "CMapLoadable_OnSetMapObjectCreateLayer",
	68: "CMapLoadable_OnSetSpineBackEffect",
	69: "CMapLoadable_OnSetSpineObjectEffect",
	70: "CMapLoadable_OnCreateSpineRectEvent",
	71: "CMapLoadable_OnRemoveSpineRectEvent",
	72: "CMapLoadable_OnSpineRE_AddBackEvent",
	73: "CMapLoadable_OnSpineRE_AddObjectEvent",
	74: "CMapLoadable_OnCreateCameraCtrlZone",
	75: "CMapLoadable_OnRemoveCameraCtrlZone",
	76: "CMapLoadable_ReloadBack",
	77: "CWvsContext_OnInventoryOperation",
	78: "CWvsContext_OnInventoryGrow",
	79: "CWvsContext_OnStatChanged",
	80: "CWvsContext_OnTemporaryStatSet",
	81: "CWvsContext_OnTemporaryStatReset",
	82: "CWvsContext_OnForcedStatSet",
	83: "CWvsContext_OnForcedStatReset",
	84: "CWvsContext_OnChangeSkillRecordResult",
	85: "CWvsContext_OnChangeStealMemoryResult",
	86: "CWvsContext_OnUserDamageOnFallingCheck",
	87: "CWvsContext_OnPersonalShopBuyCheck",
	88: "CWvsContext_OnMobDropMesoPickup",
	89: "CWvsContext_OnBreakTimeFieldEnter",
	90: "CWvsContext_OnRuneActSuccess",
	91: "CWvsContext_OnResultStealSkillList",
	92: "CWvsContext_OnSkillUseResult",
	93: "CWvsContext_OnExclRequest",
	94: "CWvsContext_OnGivePopularityResult",
	95: "CWvsContext_OnMessage",
	96: "CWvsContext_OnMemoResult",
	97: "CWvsContext_OnMapTransferResult",
	98: "CWvsContext_OnAntiMacroResult",
	99: "CWvsContext_OnAntiMacroBombResult",
	101: "CWvsContext_OnClaimResult",
	102: "CWvsContext_OnSetClaimSvrAvailableTime",
	103: "CWvsContext_OnClaimSvrStatusChanged",
	104: "CWvsContext_OnStarPlanetUserCount",
	105: "CWvsContext_OnSetTamingMobInfo",
	106: "CWvsContext_OnQuestClear",
	107: "CWvsContext_OnEntrustedShopCheckResul",
	108: "CWvsContext_OnSkillLearnItemResult",
	109: "CWvsContext_OnSkillResetItemResult",
	110: "CWvsContext_OnAbilityResetItemResult",
	111: "CWvsContext_OnExpConsumeItemResult",
	112: "CWvsContext_OnExpItemGetResult",
	113: "CWvsContext_OnCharSlotIncResult",
	114: "_",
	115: "CWvsContext_OnGatherItemResult",
	116: "CWvsContext_OnSortItemResult",
	119: "CWvsContext_OnCharacterInfo",
	120: "CWvsContext_OnPartyResult",
	121: "CWvsContext_OnPartyMemberCandidateResult",
	122: "CWvsContext_OnUrusPartyMemberCandidateResult",
	123: "CWvsContext_OnPartyCandidateResult",
	124: "CWvsContext_OnUrusPartyResult",
	125: "CWvsContext_OnIntrusionFriendCandidateResult",
	126: "CWvsContext_OnIntrusionLobbyCandidateResult",
	128: "CWvsContext_OnExpedtionResult",
	129: "CWvsContext_OnFriendResult",
	130: "CWvsContext_OnLoadAccountIDOfCharacterFriendResult",
	131: "CWvsContext_OnGuildRequest",
	132: "CWvsContext_OnGuildResult",
	133: "CWvsContext_OnAllianceResult",
	134: "CWvsContext_OnTownPortal",
	135: "CWvsContext_OnBroadcastMsg",
	136: "CAswanTimeTableManClient_OnPacket",
	137: "CWvsContext_OnIncubatorResult",
	138: "CWvsContext_OnIncubatorHotItemResult",
	139: "CWvsContext_OnShopScannerResult",
	140: "CWvsContext_OnShopLinkResult",
	141: "CWvsContext_OnAuctionResult",
	142: "CWvsContext_OnAuctionMessage",
	143: "CWvsContext_OnMarriageRequest",
	144: "CWvsContext_OnMarriageResult",
	145: "CWvsContext_OnWeddingGiftResult",
	146: "CWvsContext_OnNotifyMarriedPartnerMapTransfer",
	147: "CWvsContext_OnCashPetFoodResult",
	148: "CWvsContext_OnCashPetPickUpOnOffResult",
	149: "CWvsContext_OnCashPetSkillSettingResult",
	150: "CWvsContext_OnCashLookChangeResult",
	151: "CWvsContext_OnCashPetDyeingResult",
	152: "CWvsContext_OnSetWeekEventMessage",
	153: "CWvsContext_OnSetPotionDiscountRate",
	154: "CWvsContext_OnBridleMobCatchFail",
	155: "CWvsContext_UNK155",
	159: "CWvsContext_OnMonsterBookSetCard",
	160: "CWvsContext_OnMonsterBookSetCover",
	161: "CWvsContext_OnHourChanged",
	162: "CWvsContext_OnMiniMapOnOff",
	163: "CWvsContext_OnConsultAuthkeyUpdate",
	164: "CWvsContext_OnClassCompetitionAuthkeyUpdate",
	165: "CWvsContext_OnWebBoardAuthkeyUpdate",
	166: "CWvsContext_OnSessionValue",
	167: "CWvsContext_OnPartyValue",
	168: "CWvsContext_OnFieldSetVariable",
	169: "CWvsContext_OnFieldValue",
	170: "CWvsContext_OnBonusExpRateChanged",
	171: "CWvsContext_OnFamilyChartResult",
	172: "CWvsContext_OnFamilyInfoResult",
	173: "CWvsContext_OnFamilyResult",
	174: "CWvsContext_OnFamilyJoinRequest",
	175: "CWvsContext_OnFamilyJoinRequestResult",
	176: "CWvsContext_OnFamilyJoinAccepted",
	177: "CWvsContext_OnFamilyPrivilegeList",
	178: "CWvsContext_OnFamilyFamousPointIncResult",
	179: "CWvsContext_OnFamilyNotifyLoginOrLogout",
	180: "CWvsContext_OnFamilySetPrivilege",
	181: "CWvsContext_OnFamilySummonRequest",
	182: "CWvsContext_OnNotifyLevelUp",
	183: "CWvsContext_OnNotifyWedding",
	184: "CWvsContext_OnNotifyJobChange",
	185: "CWvsContext_OnSetBuyEquipExt",
	186: "CWvsContext_OnSetPassenserRequest",
	187: "CWvsContext_OnScriptProgressMessageBySoul",
	188: "CWvsContext_OnScriptProgressMessage",
	189: "CWvsContext_OnScriptProgressItemMessage",
	190: "CWvsContext_OnStaticScreenMessage",
	191: "CWvsContext_OffStaticScreenMessage",
	192: "CWvsContext_OnWeatherEffectNotice",
	193: "CWvsContext_OnWeatherEffectNoticeY",
	194: "CWvsContext_OnProgressMessageFont",
	195: "CWvsContext_OnDataCRCCheckFailed",
	196: "CWvsContext_OnShowSlotMessage",
	197: "CWvsContext_OnWildHunterInfo",
	198: "CWvsContext_OnZeroInfo",
	199: "CWvsContext_OnZeroWP",
	200: "CWvsContext_OnZeroInfoSubHP",
	201: "_",
	202: "CWvsContext_ClearAnnouncedQuest",
	203: "CWvsContext_OnResultInstanceTable",
	204: "CWvsContext_OnCoolTimeSet",
	205: "CWvsContext_OnItemPotChange",
	206: "CWvsContext_OnSetItemCoolTime",
	207: "CWvsContext_OnSetAdDisplayInfo",
	208: "CWvsContext_OnSetAdDisplayStatus",
	209: "CWvsContext_OnSetSonOfLinkedSkillResult",
	210: "CWvsContext_OnSetMapleStyleInfo",
	211: "CWvsContext_OnSetBuyLimitCount",
	212: "CWvsContext_OnResetBuyLimitCount",
	213: "CWvsContext_OnUpdateUIEventListInfo",
	214: "_",
	215: "CUIFieldItem_OnPacket",
	216: "CUIFieldItemInventory_OnPacket",
	218: "CWvsContext_OnResultSetStealSkill",
	219: "CWvsContext_OnSlashCommand",
	220: "CWvsContext_OnStartNavigation",
	221: "CWvsContext_OnFunckeySetByScript",
	222: "CWvsContext_OnCharacterPotentialSet",
	223: "CWvsContext_OnCharacterPotentialReset",
	224: "CWvsContext_OnCharacterHonorExp",
	225: "_",
	226: "_",
	227: "CWvsContext_OnReadyForRespawn",
	228: "CWvsContext_OnReadyForRespawnByPoint",
	229: "CWvsContext_OpenReadyForRespawnUI",
	230: "CWvsContext_OnCharacterHonorGift",
	231: "CWvsContext_OnCrossHunterCompleteResult",
	232: "CWvsContext_OnCrossHunterShopResult",
	233: "CWvsContext_OnSetCashItemNotice",
	234: "CWvsContext_OnSetSpecialCashItem",
	235: "CWvsContext_OnShowEventNotice",
	236: "CWvsContext_OnBoardGameResult",
	237: "_",
	238: "CWvsContext_OnValuePackResult",
	239: "CWvsContext_OnNaviFlyingResult",
	240: "CWvsContext_OnMapleStyleResult",
	241: "CWvsContext_OnCheckWeddingExResult",
	242: "CWvsContext_OnBingoResult",
	243: "CWvsContext_OnBingoCassandraResult",
	244: "CWvsContext_OnUpdateVIPGrade",
	245: "CWvsContext_OnMesoRangerResult",
	246: "CWvsContext_OnSetMaplePoint",
	247: "CWvsContext_OnSetAdditionalCashInfo",
	248: "CWvsContext_OnSetMiracleTimeInfo",
	249: "CWvsContext_OnHyperSkillResetResult",
	250: "CWvsContext_OnGetServerTime",
	251: "CWvsContext_OnGetCharacterPosition",
	253: "CWvsContext_OnReturnEffectConfirm",
	254: "CWvsContext_OnReturnEffectModified",
	255: "CWvsContext_OnWhiteAddtionalCubeResult",
	256: "CWvsContext_OnBlackCubeResult",
	257: "CWvsContext_OnMemorialCubeResult",
	258: "CWvsContext_OnMemorialCubeModified",
	259: "CWvsContext_OnDressUpInfoModified",
	260: "CWvsContext_OnResetOnStateForOnOffSkill",
	261: "_",
	262: "CWvsContext_OnSetOffStateForOnOffSkill",
	263: "CWvsContext_OnIssueReloginCookie",
	264: "CWvsContext_OnAvatarPackTest",
	265: "CWvsContext_OnEvolvingResult",
	266: "CWvsContext_OnActionBarResult",
	267: "CUIGuildContentRank_OnPacket",
	268: "CWvsContext_OnGuildSearchResult",
	269: "CUIButterFlyCustomize_OnPacket",
	270: "CWvsContext_OnHalloweenCandyRankingResult",
	271: "CWvsContext_OnGetRewardResult",
	272: "CWvsContext_OnMentoring",
	273: "CWvsContext_OnGetLotteryResult",
	274: "CWvsContext_OnCheckProcessResult",
	275: "CWvsContext_OnCompleteNpcSpeechSuccess",
	276: "CWvsContext_OnCompleteSpecialCheckSuccess",
	277: "CWvsContext_OnSetAccountInfo",
	278: "CWvsContext_OnSetGachaponFeverTimeInfo",
	279: "CWvsContext_OnAvatarMegaphoneRes",
	280: "CWvsContext_OnSetAvatarMegaphone",
	281: "CWvsContext_OnClearAvatarMegaphone",
	282: "CWvsContext_OnRequestEventList",
	283: "CWvsContext_OnLikePoint",
	284: "CWvsContext_OnSignErrorAck",
	285: "CWvsContext_OnAskAfterErrorAck",
	286: "CWvsContext_OnEventNameTag",
	287: "CWvsContext_OnAcquireEventNameTag",
	288: "CWvsContext_OnJobFreeChangeResult",
	289: "CWvsContext_OnEventLotteryOpen",
	290: "CWvsContext_OnEventLotteryResult",
	291: "CInvasionSupportMan_OnInvasionSupportStateChange",
	292: "CInvasionSupportMan_OnInvasionSupportAttackResult",
	293: "CInvasionSupportMan_OnBossKillResult",
	294: "CInvasionSupportMan_OnInvasionSupportSettingResult",
	295: "CInvasionSupportMan_OnInvasionElapsedTime",
	296: "CInvasionSupportMan_OnInvasionSystemMsg",
	298: "CWvsContext_OnScreenMsg",
	299: "CWvsContext_OnTradeBlockForSnapShot",
	300: "CWvsContext_OnLimitGoodsNoticeResult",
	301: "CWvsContext_OnMonsterBattleSystemResult",
	302: "CWvsContext_OnMonsterBattleCombatResult",
	303: "CWvsContext_OnUniverseBossPossible",
	304: "CWvsContext_OnUniverseBossImpossible",
	305: "CWvsContext_OnCashShopPreviewInfo",
	306: "CWvsContext_OnChangeSoulCollectionResult",
	307: "CWvsContext_OnSelectSoulCollectionResult",
	308: "CWvsContext_OnMasterPieceReward",
	309: "CWvsContext_OnPendantSlotIncResult",
	310: "CWvsContext_OnBossArenaMatchSuccess",
	311: "CWvsContext_OnBossArenaMatchFail",
	312: "CWvsContext_OnBossArenaMatchRequestDone",
	313: "CWvsContext_OnUserSoulMatching",
	314: "CWvsContext_OnCatapultUpgradeSkill",
	315: "CWvsContext_OnCatapultResetSkill",
	316: "CWvsContext_OnPartyQuestRankingResult",
	317: "CWvsContext_OnSetCoordinationContestInfo",
	318: "CWvsContext_OnWorldTransferResult",
	319: "CWvsContext_OnTrunkSlotIncResult",
	320: "CWvsContext_OnEliteMobWMI",
	321: "CWvsContext_OnRandomPortalNotice",
	322: "CWvsContext_OnNotifyWorldTransferHelper",
	323: "CUIEquipmentEnchant_OnPacket",
	326: "CWvsContext_OnTopTowerRankResult",
	327: "CWvsContext_OnFriendTowerRankResult",
	328: "CWvsContext_OnTowerResultUIOpen",
	329: "CWvsContext_CKoreanJumpingGame",
	330: "CWvsContext_OnCreateSwingGame",
	331: "CWvsContext_OnUpdateMapleTVShowTime",
	332: "CWvsContext_OnReturnToTitle",
	333: "CWvsContext_OnReturnToCharacterSelect",
	334: "CWvsContext_OnFlameWizardFlameWalkEffect",
	335: "CWvsContext_OnFlameWizardFlareBlink",
	336: "CWvsContext_OnSummonedAvatarSync",
	337: "CWvsContext_OnCashShopEventInfo",
	338: "CWvsContext_OnBlackList",
	339: "CWvsContext_OnOpenUITest",
	340: "CWvsContext_OnBlackListView",
	341: "CWvsContext_OnScrollUpgradeFeverTime",
	342: "CWvsContext_OnTextEquipInfo",
	343: "CWvsContext_OnTextEquipUIOpen",
	344: "CWvsContext_OnUIStarPlanetMiniGameResult",
	345: "CWvsContext_OnUIStarPlanetTrendShop",
	346: "CWvsContext_OnUIStarPlanetMiniGameQueue",
	348: "CWvsContext_OnStarPlanetRoundInfo",
	349: "CWvsContext_OnStarPlanetResult",
	350: "CWvsContext_OnBackSpeedCtrl",
	351: "CWvsContext_OnSetMazeArea",
	352: "CWvsContext_OnCharacterBurning",
	353: "CWvsContext_OnBattleStatCoreInfo",
	354: "CWvsContext_OnBattleStatCoreAck",
	355: "CWvsContext_OnGachaponTestResult",
	356: "CWvsContext_OnMasterPieceTestResult",
	357: "CWvsContext_OnRoyalStyleTestResult",
	358: "CWvsContext_OnBeautyCouponTestResult",
	359: "CWvsContext_OnNickSkillExpired",
	360: "CWvsContext_OnRandomMissionResult",
	361: "CWvsContext_On12thTresureResult",
	362: "CWvsContext_On12thTresureBuff",
	363: "CWvsContext_OnItemCollectionResult",
	364: "CWvsContext_OnCheckCollectionCompleteResult",
	365: "CWvsContext_OnItemCollectionList",
	366: "CWvsContext_OnReceiveToadsHammerRequestResult",
	367: "CWvsContext_OnReceiveHyperStatSkillResetResult",
	368: "CWvsContext_OnInventoryOperationResult",
	369: "CWvsContext_OnGetSavedUrusSkill",
	370: "CWvsContext_OnSetRolePlayingCharacterInfo",
	371: "CWvsContext_OnMVPAlarm",
	372: "CWvsContext_OnMonsterCollectionResult",
	373: "CWvsContext_OnTowerChairSettingResult",
	374: "CWvsContext_OnNeedClientResponse",
	375: "CWvsContext_OnCharacterModified",
	376: "_",
	377: "CWvsContext_OnTradeKingShopItem",
	378: "CWvsContext_OnTradeKingShopRes",
	379: "CWvsContext_OnPlatFormarEnterResult",
	380: "CWvsContext_OnPlatFormarOxyzen",
	381: "CWvsContext_UNKStart",
	388: "CWvsContext_OnHairStyleCoupon",
	399: "CWvsContext_UNK399",
	411: "CWvsContext_UNK411",
	441: "CWvsContext_UNK441",
	442: "CStage_OnSetField",
	443: "CStage_OnSetFarmField",
	444: "CStage_OnSetCashShop",
	445: "CStage_UNK445",
	446: "CField_OnTransferFieldReqIgnored",
	447: "CField_OnTransferChannelReqIgnored",
	448: "CField_OnTransferPvpReqIgnored",
	449: "CField_OnFieldSpecificData",
	450: "CField_OnGroupMessage",
	451: "CField_OnFieldUniverseMessage",
	452: "CField_OnWhisper",
	453: "CField_OnSummonItemInavailable",
	454: "CField_OnFieldEffect",
	455: "CField_OnBlowWeather",
	456: "CField_OnPlayJukeBox",
	457: "CField_OnAdminResult",
	458: "CField_OnQuiz",
	459: "CField_OnDesc",
	460: "_",
	463: "CField_OnSetQuestClear",
	464: "CField_OnSetQuestTime",
	465: "CField_OnSetObjectState",
	466: "CWnd_OnDestroy",
	468: "CField_OnStalkResult",
	471: "CQuickslotKeyMappedMan_OnInit",
	472: "CField_OnFootHoldMove",
	473: "CField_OnCorrectFootHoldMove",
	474: "_",
	475: "_",
	476: "_",
	478: "CField_OnSmartMobNotice",
	479: "CField_OnChangePhase",
	480: "CField_OnChangeMobZone",
	482: "CField_OnPvPMigrateInfoResult",
	483: "CField_OnCurNodeEventEnd",
	484: "CField_OnCreateForceAtom",
	485: "CField_OnSetAchieveRate",
	486: "CField_OnSetQuickMoveInfo",
	487: "CField_OnChangeAswanSiegeWeaponGauge",
	488: "CField_OnCreateObtacle",
	489: "CField_OnClearObtacle",
	490: "_",
	491: "CField_OnB2FootHoldCreate",
	492: "CField_OnDebuffObjON",
	493: "_",
	494: "_",
	495: "CField_OnCreateFallingCatcher",
	496: "CField_OnChaseEffectSet",
	497: "CField_OnSetMirrorDungeonInfo",
	498: "CField_OnSetIntrusion",
	499: "CField_OnCannotDrop",
	500: "CField_OnFootHoldOnOf",
	501: "CField_OnLadderRopeOnOff",
	502: "CField_OnMomentAreaOnOff",
	503: "CField_OnMomentAreaOnOffAll",
	504: "CField_OnChatLetClientConnect",
	505: "CUICoordinationContest_OnPacket",
	506: "CField_OnEliteState",
	507: "CField_OnPlaySound",
	508: "CField_OnStackEventGauge",
	509: "CField_OnSetUnionField",
	511: "CField_OnStarPlanetBurningTimeInfo",
	512: "CField_OnPublicShareState",
	513: "CField_OnFunctionTempBlock",
	514: "CUIStatusBar_OnPacket",
	515: "FieldDelaySkill_OnFieldSkillDelay",
	516: "CField_OnWeatherPacket_Add",
	517: "CField_OnWeatherPacket_Remove",
	518: "CField_OnWeatherPacket_Msg",
	519: "CField_OnAddWreckage",
	520: "CField_OnDelWreckage",
	522: "CField_OnCreateMirrorImage",
	523: "CField_OnFuntionFootholdMan",
	524: "_",
	529: "CField_OnMouseMove",
	530: "CField_UNK530",
	531: "CUserPool_OnUserEnterField",
	532: "CUserPool_OnUserLeaveField",
	533: "CUser_OnChat",
	534: "CUser_OnADBoard",
	535: "CUser_OnMiniRoomBalloon",
	536: "CUser_SetConsumeItemEffect",
	537: "CUser_ShowItemUpgradeEffect",
	539: "CUser_ShowItemSkillSocketUpgradeEffect",
	540: "CUser_ShowItemSkillOptionUpgradeEffect",
	541: "CUser_ShowItemReleaseEffect",
	542: "CUser_ShowItemUnreleaseEffect",
	543: "CUser_ShowItemLuckyItemEffect",
	544: "CUser_ShowItemMemorialEffect",
	545: "CUser_ShowItemAdditionalUnReleaseEffect",
	546: "CUser_ShowItemAdditionalSlotExtendEffect",
	547: "CUser_ShowItemFireWorksEffect",
	548: "CUser_ShowItemOptionChangeEffect",
	549: "CUser_OnRedCubeResult",
	550: "_",
	551: "CUser_OnHitByUser",
	552: "CUser_OnDotByUser",
	553: "CUser_OnResetAllDot",
	554: "CUser_OnDamageByUser",
	555: "CUser_OnTeslaTriangle",
	556: "CUser_OnFollowCharacter",
	557: "CUser_OnShowPQReward",
	558: "CUser_OnSetOneTimeAction",
	559: "CUser_OnMakingSkillResult",
	560: "CUser_OnSetMakingMeisterSkillEff",
	561: "CUser_OnGatherResult",
	562: "CUser_OnUserExplode",
	563: "CUser_OnUserHitByCounter",
	564: "CUser_OnPyramidLethalAttack",
	565: "CUser_OnMixerResult",
	566: "CUser_OnWaitQueueResponse",
	567: "CUser_OnCategoryEventNameTag",
	568: "CUser_OnSetDamageSkin",
	569: "CUser_OnSetPremiumDamageSkin",
	570: "CUser_OnSetSoulEffect",
	571: "CUser_OnSitResult",
	572: "_",
	573: "CUser_OnStarPlanetPointInfo",
	574: "CUser_OnStarPlanetAvatarLookSet",
	575: "CUser_OnTossedByMobSkill",
	576: "CUser_OnBattleAttackHit",
	577: "CUser_OnBattleUserHitByMob",
	578: "CUser_OnFreezeHotEventInfo",
	579: "CUser_OnEventBestFriendInfo",
	580: "CUser_OnSetRepeatOneTimeAction",
	581: "CUser_OnSetReplaceMoveAction",
	582: "CUser_OnInGameCubeResult",
	583: "_",
	584: "CUser_OnSetActiveEmoticonItem",
	585: "CUser_OnCreatePsychicLock",
	586: "CUser_OnRecreatePathPsychicLcok",
	587: "CUser_OnReleasePsychicLock",
	588: "CUser_OnReleasePsychicLockMob",
	589: "CUser_OnCreatePsychicArea",
	590: "CUser_OnReleasePsychicArea",
	591: "CUser_OnRWZeroBunkerMobBind",
	592: "CUser_OnBeastFormWingOnOff",
	593: "CUser_OnSetMesoChairCount",
	594: "CUser_OnRefreshNameTagMark",
	595: "CUser_OnStigmaEffect",
	596: "CUser_UNKStart",
	608: "CUser_UNK608",
	609: "CPet_OnActivated",
	610: "CPet_OnMove",
	611: "CPet_OnAction",
	612: "CPet_OnActionSpeak",
	613: "CPet_OnNameChanged",
	614: "CPet_OnLoadExceptionList",
	615: "CPet_OnHueChanged",
	616: "CPet_OnModified",
	619: "CPet_OnActionCommand",
	620: "CDragon_OnCreated",
	621: "CDragon_OnMove",
	622: "CDragon_ReMove",
	623: "CDragon_ReMoveBack",
	624: "CAndroid_OnCreated",
	625: "CAndroid_OnMove",
	626: "CAndroid_OnActionSet",
	627: "CAndroid_OnModified",
	628: "CAndroid_ReMove",
	629: "CFoxMan_EnterField",
	630: "CFoxMan_OnMove",
	631: "CFoxMan_OnExclResult",
	632: "CFoxMan_OnShowChangeEffect",
	633: "CFoxMan_OnModified",
	634: "CFoxMan_LeaveField",
	636: "CSkillPet_OnMove",
	637: "CSkillPet_OnAction",
	638: "CSkillPet_OnState",
	641: "CSkill_UNK",
	651: "CUserRemote_OnAttack",
	652: "CUserRemote_OnSkillPrepare",
	653: "CUserRemote_OnMovingShootAttackPrepare",
	654: "CUserRemote_OnSkillCancel",
	655: "CUserRemote_OnHit",
	656: "CUserRemote_OnEmotion",
	657: "CUserRemote_OnAndroidEmotion",
	658: "CUserRemote_OnSetActiveEffectItem",
	659: "CUserRemote_OnSetMonkeyEffectItem",
	660: "CUserRemote_OnSetActiveNickItem",
	661: "CUserRemote_OnSetDefaultWingItem",
	662: "CUserRemote_OnSetKaiserTransformItem",
	664: "CUserRemote_OnShowUpgradeTombEffect",
	665: "CUserRemote_OnSetActivePortableChair",
	666: "CUserRemote_OnAvatarModified",
	667: "CUserRemote_OnEffect",
	668: "CUserRemote_OnSetTemporaryStat",
	669: "CUserRemote_OnResetTemporaryStat",
	670: "CUserRemote_OnReceiveHP",
	671: "CUserRemote_OnGuildNameChanged",
	672: "CUserRemote_OnGuildMarkChanged",
	673: "CUserRemote_OnPvPTeamChanged",
	674: "CUserRemote_OnGatherActionSet",
	675: "CUserRemote_OnUpdatePvPHPTag",
	676: "CUserRemote_OnDragonGlide",
	677: "CUserRemote_OnKeyDownAreaMovePath",
	678: "CUserRemote_OnLaserInfoForRemote",
	679: "CUserRemote_OnKaiserColorOrMorphChange",
	680: "CUserRemote_OnDestroyGrenade",
	681: "CUser_OnSetItemAction",
	682: "CUserRemote_OnZeroTag",
	683: "CUserRemote_OnIntrusion",
	684: "CUserRemote_OnZeroLastAssistState",
	685: "CUserRemote_OnSetMoveGrenade",
	686: "CUserRemote_OnSetCustomizeEffect",
	687: "CUserRemote_OnRuneStoneAction",
	688: "CUserRemote_OnKinesisPsychicEnergyShieldEffect",
	689: "CUserRemote_OnDragonAction",
	690: "CUserRemote_OnDragonBreathEarthEffect",
	691: "CUserRemote_OnReleaseRWGrab",
	692: "CUserRemote_OnRWMultiChargeCancelRequest",
	693: "CUserRemote_OnScouterMaxDamageUpdate",
	694: "CUserRemote_OnStigmaDeliveryResponse",
	701: "CUserRemote_OnThrowGrenade",
	702: "CUserLocal_OnEmotion",
	703: "CUserLocal_OnAndoridEmotion",
	704: "CUserLocal_OnEffect",
	705: "CUserLocal_OnTeleport",
	707: "CUserLocal_OnMesoGive_Succeeded",
	708: "CUserLocal_OnMesoGive_Failed",
	709: "CUserLocal_OnQuestResult",
	710: "CUserLocal_OnNotifyHPDecByField",
	711: "CUserLocal_OnUserPetSkillChanged",
	712: "CUserLocal_OnBalloonMsg",
	713: "CUserLocal_OnPlayEventSound",
	714: "CUserLocal_OnPlayMinigameSound",
	715: "CUserLocal_OnMakerResult",
	716: "CUserLocal_OnOpenConsultBoard",
	717: "CUserLocal_OnOpenClassCompetitionPage",
	718: "CUserLocal_OnOpenUI",
	719: "CUserLocal_OnCloseUI",
	720: "CUserLocal_OnOpenUIWithOption",
	721: "CUserLocal_OnOpenWebUI",
	722: "CUserLocal_OnSetDirectionMode",
	723: "CUserLocal_OnSetInGameDirectionMode",
	724: "CUserLocal_OnSetStandAloneMode",
	725: "CUserLocal_OnHireTutor",
	726: "CUserLocal_OnTutorMsg",
	727: "CUserLocal_OnHireTutorById",
	728: "CUserLocal_OnSetPartner",
	729: "CUserLocal_OnSetPartnerAction",
	730: "CUserLocal_OnSetPartnerForceFlip",
	731: "CUserLocal_OnSwitchRP",
	732: "CUserLocal_OnModComboResponse",
	733: "CUserLocal_OnIncComboResponseByComboRecharge",
	734: "CUserLocal_OnRadioSchedule",
	735: "CUserLocal_OnOpenSkillGuide",
	736: "CUserLocal_OnNoticeMsg",
	737: "CUserLocal_OnChatMsg",
	738: "CUserLocal_OnSetUtilDlg",
	739: "CUserLocal_OnBuffzoneEffect",
	740: "CUserLocal_OnTimeBombAttack",
	741: "CUserLocal_OnExplosionAttack",
	742: "CUserLocal_OnPassiveMove",
	743: "CUserLocal_OnFollowCharacterFailed",
	744: "CUserLocal_SetNextShootExJablin",
	745: "CUICreatePremiumAdventurer_OnResult",
	746: "CUserLocal_OnGatherRequestResult",
	747: "CUserLocal_OnRuneStoneUseAck",
	748: "CUserLocal_OnBagItemUseResult",
	749: "CUserLocal_OnRandomTeleportKey",
	750: "CUserLocal_OnSetGagePoint",
	751: "CInGameDirectionEvent_OnInGameDirectionEvent",
	752: "CUserLocal_OnMedalReissueResult",
	753: "CUserLocal_OnDodgeSkillReady",
	754: "CUserLocal_OnRemoveMicroBuffSkill",
	755: "CUserLocal_OnVideoByScript",
	756: "_",
	757: "CUserLocal_OnRewardMobListResult",
	758: "CUserLocal_OnIncJudgementStackResponse",
	759: "CUserLocal_OnIncCharmByCashPRMsg",
	760: "CUserLocal_OnSetBuffProtector",
	761: "CUserLocal_OnIncLarknessResponse",
	762: "CUserLocal_OnDetonateBomb",
	763: "CUserLocal_OnAggroRankInfoName",
	764: "CUserLocal_OnDeathCountInfo",
	765: "CUserLocal_OnDeathCountInfo2",
	767: "CUserLocal_OnSeverAckMobZoneStateChange",
	770: "CUserLocal_OnRandomEmotion",
	771: "CUserLocal_OnSetFlipTheCoinEnabled",
	772: "CUserLocal_OnTrickOrTreatResult",
	773: "CUserLocal_OnGiantPetBuff",
	774: "CUserLocal_OnB2BodyResult",
	775: "CUserLocal_OnSetDead",
	776: "CUserLocal_OnOpenUIOnDead",
	777: "CUserLocal_OnExpiredNotice",
	778: "CUserLocal_OnDoLotteryUI",
	779: "CUserLocal_OnRouletteStart",
	780: "CUserLocal_OnSitOnTimeCapsule",
	781: "CUserLocal_OnSitOnDummyPortableChair",
	785: "CUserLocal_OnFinalAttackRequest",
	786: "CUserLocal_OnSetGun",
	787: "CUserLocal_OnSetAmmo",
	788: "CUserLocal_OnCreateGun",
	789: "CUserLocal_OnClearGun",
	790: "CUserLocal_OnResultShootAttackInFPSMode",
	791: "CUserLocal_OnMirrorDungeonEnterFail",
	792: "CUserLocal_OnMirrorDungeonUnitCleared",
	793: "CWvsContext_RegisterMirrorDungeonBoss",
	794: "CUserLocal_OnMirrorDungeonRecord",
	795: "CUserLocal_OnOpenURL",
	796: "CUserLocal_OnZeroCombatRecovery",
	797: "CUserLocal_OnMirrorStudyUIOpen",
	798: "CUserLocal_OnSkillCooltimeReduce",
	799: "CUserLocal_OnMirrorReadingUIOpen",
	800: "CUserLocal_OnUserCtrlMobSkill_QPush",
	801: "CUserLocal_OnZeroLevelUpAlram",
	802: "CUserLocal_OnUserCtrlMobSkill_QPop",
	803: "CUserLocal_OnUserCtrlMobSkill_Fail",
	804: "CUserLocal_OnForceSummonedRemove",
	805: "CUserLocal_OnUserRespawn",
	806: "CUserLocal_OnUserCtrlMobSkill_ForcedPop",
	807: "CUserLocal_OnMonsterBattleCapture",
	808: "CUserLocal_OnIsUniverse",
	809: "CUserLocal_OnPortalGroup",
	810: "CUserLocal_OnSetMovable",
	811: "CUserLocal_OnUserCtrlMobSkill_PushCoolTime",
	812: "CUserLocal_OnMoveParticleEff",
	813: "CUserLocal_OnDoActiveEventSkillByScript",
	814: "CUserLocal_OnSetStatusbarJobNameBlur",
	815: "CUserLocal_OnRuneStoneSkillAck",
	816: "CUserLocal_ResetRuneStoneAction",
	817: "CUserLocal_OnMoveToContentsCannotMigrate",
	818: "CUserLocal_OnPlayAmbientSound",
	819: "CUserLocal_OnPlayAmbientSound2",
	820: "CUserLocal_OnStopAmbientSound",
	821: "CUserLocal_OnFlameWizardElementFlameSummon",
	822: "CUserLocal_OnCameraMode",
	823: "CUserLocal_OnSpotlightToCharacter",
	824: "CUserLocal_OnBossPartyCheckDone",
	825: "CUserLocal_OnFreeLookChangeUIOpen",
	826: "CUserLocal_OnFreeLookChangeSuccess",
	827: "CUserLocal_OnGrayBackground",
	828: "CUserLocal_OnGetNpcCurrentAction",
	829: "CUserLocal_OnCameraRotation",
	830: "CUserLocal_OnCameraSwitch",
	831: "CameraCtrl_Network_OnPacket",
	832: "CUserLocal_OnUserSetFieldFloating",
	833: "CUserLocal_OnAddPopupSay",
	834: "CPopupSayMan_RemovePopupSay",
	835: "CUserLocal_OnJaguarSkill",
	836: "CUserLocal_OnActionLayerRelmove",
	837: "CUserLocal_OnClientResolution",
	838: "CUserLocal_OnUserBonusAttackRequest",
	839: "CUserLocal_OnUserRandAreaAttackRequest",
	840: "CUserLocal_OnJaguarActive",
	841: "CUserLocal_OnSkillCooltimeSetM",
	842: "CUserLocal_OnSetCarryReactorInfo",
	843: "CUserLocal_OnReactorSkillUseRequest",
	844: "CUserLocal_OnOpenBattlePvPChampSelectUI",
	845: "CUserLocal_OnBattlePvPItemDropSound",
	846: "CUserLocal_OnSetMesoCountByScript",
	847: "CUserLocal_OnPlantPotClickResult",
	848: "CUserLocal_OnPlantPotEffect",
	849: "CUserLocal_OnDamage",
	850: "CUserLocal_OnRoyalGuardAttack",
	851: "CUserLocal_OnDoActivePsychicArea",
	852: "CUserLocal_OnEnterFieldPsychicInfo",
	853: "CUserLocal_OnLeaveFieldPsychicInfo",
	854: "CUserLocal_OnTouchMeStateResult",
	855: "CUserLocal_OnFieldScoreUpdate",
	856: "CUserLocal_OnUrusReusltUIOpen",
	857: "CUserLocal_SetNoMoreLife",
	858: "CUserLocal_UNK858",
	859: "CUserLocal_OnCreateAreaDotInfo",
	860: "CUserLocal_OnSetSlowDown",
	861: "CUserLocal_OnRegisterExtraSkill",
	862: "CUserLocal_OnResWarriorLiftMobInfo",
	863: "CUserLocal_OnUserRenameResult",
	864: "CUserLocal_OnDamageSkinSaveResult",
	865: "CUserLocal_OnStigmaRemainTime",
	866: "CUserLocal_OnOpenMesoSackSuccess",
	867: "CUserLocal_OnOpenMesoSackFail",
	868: "CUserLocal_UNKStart",
	883: "CUserLocal_UNK883",
	916: "CUserLocal_UNKEnd",
	917: "CUserLocal_OnSkillCooltimeSet",
	919: "CSummonedPool_OnCreated",
	920: "CSummonedPool_OnRemoved",
	921: "CSummonedPool_OnMove",
	922: "CSummonedPool_OnAttack",
	923: "CSummonedPool_OnAttackPvP",
	924: "CSummonedPool_OnSetReference",
	925: "CSummonedPool_OnSkill",
	926: "CSummonedPool_OnSkillPvP",
	927: "CSummonedPool_OnUpdateHPTag",
	928: "CSummonedPool_OnAttackDone",
	929: "CSummonedPool_OnSetResist",
	930: "CSummonedPool_OnSummonedActionChange",
	931: "CSummonedPool_OnAssistAttackRequest",
	932: "CSummonedPool_OnSummonAttackActive",
	933: "CSummonedPool_OnSummonBeholderRevengeAttack",
	934: "CSummonedPool_OnHit",
	935: "CMobPool_OnMobEnterField",
	936: "CMobPool_OnMobLeaveField",
	937: "CMobPool_OnMobChangeController",
	938: "CMobPool_OnMobSetAfterAttack",
	939: "CMobPool_OnMobBlockAttack",
	940: "CMobPool_UNKStart",
	950: "CMobPool_UNKEnd",
	951: "CMobPool_OnMobCrcKeyChanged",
	941: "CMob_OnMove",
	942: "CMob_OnCtrlAck",
	944: "CMob_OnStatSet",
	945: "CMob_OnStatReset",
	946: "CMob_OnSuspendReset",
	947: "CMob_OnAffected",
	948: "CMob_OnDamaged",
	949: "CMob_OnSpecialEffectBySkill",
	952: "CMob_OnCrcDataRequest",
	953: "CMob_OnHPIndicator",
	954: "CMob_OnCatchEffect",
	955: "CMob_OnStealEffect",
	956: "CMob_OnEffectByItem",
	957: "CMob_OnMobSpeaking",
	958: "CMob_OnMobMessaging",
	959: "CMob_OnMobSkillDelay",
	960: "CMob_OnEscortFullPath",
	961: "CMob_OnEscortStopEndPermmision",
	962: "CMob_OnEscortStopByScript",
	963: "CMob_OnEscortStopSay",
	964: "CMob_OnEscortReturnBefore",
	965: "CMob_OnNextAttack",
	966: "CMob_OnMobTeleportRequest",
	967: "CMob_OnForcedAction",
	968: "CMob_OnForcedSkillAction",
	969: "CMob_UNK969",
	970: "CMob_OnTimeResist",
	971: "CMob_OnOnekillDamage",
	972: "CMob_OnAttackBlock",
	973: "CMob_OnAttackPriority",
	974: "OnAttackTimeInfo",
	975: "CMob_OnDamageShareInfoToLocal",
	976: "CMob_OnDamageShareInfoToRemote",
	977: "CMob_OnBreakDownTimeZoneTimeOut",
	978: "CMob_OnMoveAreaSet",
	979: "CMob_OnDoSkillByHit",
	980: "CMob_OnCastingBarSkill",
	981: "CMob_OnFlyTarget",
	982: "CMob_OnBounceAttackSkill",
	983: "CMob_OnAreaInstallByHit",
	984: "CMob_OnLtrbDamageSkill",
	985: "CMob_OnSummonSubBody",
	986: "CMob_OnLaserControl",
	987: "CMob_OnScale",
	988: "_",
	989: "_",
	990: "CMob_ForceChase",
	991: "CMob_OnHangOverRequest",
	992: "CMob_OnHangOverReleaseRequest",
	993: "CMob_OnDeadFPSMode",
	994: "CMob_OnAirHit",
	995: "CMob_OnDemianDelayedAttackCreate",
	996: "CMob_OnRegisterRelMobZone",
	997: "CMob_OnUnregisterRelMobZone",
	998: "CMob_OnNextTargetFromSvr",
	999: "CMob_UNKStart",
	1005: "CMob_UNKEND",
	1006: "CMob_OnMobAttackedByMob",
	156: "CNpcPool_OnNpcImitateData",
	158: "CNpcPool_OnUpdateLimitedDisableInfo",
	1012: "CNpcPool_OnNpcEnterField",
	1013: "CNpcPool_OnNpcLeaveField",
	1014: "CNpcPool_OnNpcEnterFieldForQuickMove",
	1015: "CNpcPool_OnNpcChangeController",
	1016: "CNpc_OnMove",
	1017: "CNpc_OnUpdateLimitedInfo",
	1018: "CNpc_OnSetQuizScore",
	1019: "CNpc_OnSetQuizScoreAni",
	1020: "CNpc_OnSetForceMove",
	1021: "CNpc_OnSetForceFlip",
	1022: "CNpc_UNK1022",
	1023: "CNpc_OnSetEmotion",
	1024: "CNpc_OnSetCharacterBaseAction",
	1025: "CNpc_OnViewOrHide",
	1026: "CNpc_OnPresentItemSet",
	1027: "CNpc_OnPresentTimeSet",
	1028: "CNpc_OnResetSpecialAction",
	1029: "CNpc_OnSetScreenInfo",
	1030: "CNpc_OnLocalNpcRepeatEffect",
	1031: "CNpc_OnSetNoticeBoardInfo",
	1033: "CNpc_UNK1033",
	1034: "CNpc_OnSetSpecialAction",
	1221: "CField_TypingGame_OnWaveInfo",
	1222: "CField_TypingGame_OnEnter",
	1223: "CField_TypingGame_OnMoveToPortal",
	1224: "CField_TypingGame_OnStart",
	1225: "CField_TypingGame_OnResetWordMob",
	1226: "CField_TypingGame_OnScoreInfo",
	1227: "CField_TypingGame_OnResultInfo",
	1228: "CField_TypingGame_OnTypingEffect",
	1229: "CField_TypingGame_OnFieldState",
	1230: "CField_TypingGame_OnPointEffect",
	1231: "CField_TypingGame_OnUserSkillAttackResult",
	1232: "CField_TypingGame_OnSummonWord",
	1323: "CScriptMan_OnScriptMessage",
	1393: "CCashShop_OnChargeParamResult",
	1394: "CCashShop_OnQueryCashResult",
	1395: "CCashShop_OnCashItemResult",
	1396: "CCashShop_OnPurchaseExpChanged",
	1397: "CCashShop_OnGiftMateInfoResult",
	1398: "CCashShop_OnCashShopCharStatChanged",
	1399: "CCashShop_OnCashShopTerminated",
	1400: "CCashShop_OnRedrawList",
	1401: "CCashShop_OnCoodinationResult",
	1402: "CCashShop_OnScreenMsg_",
	1403: "CCashShop_OnCash_MVP_ItemGive_Result",
	1404: "CCashShop_OnCheckMileageResult",
	1405: "CCashShop_OnChargeMileageNotice",
	1407: "CCashShop_UNKStart",
	1422: "CCashShop_UNKEnd",
	1423: "CCashShop_OnMemberShopResult",
	1425: "CUIContext_OnAttendanceEventData",
	1426: "CUIContext_OnAttendanceEventUIData",
	1427: "CUIContext_OnEventUIData",
	1428: "CUIContext_OnEventUIAck",
	1429: "CUIContext_OnGhostPaintsPointUpdate",
	1430: "CUIUserTimerMan_OnPacketProcess",
	1432: "CUIContext_OnSADResultUIOpen",
	1433: "CUIContext_OnSADLevelUpEquipDiff",
	1434: "CUIContext_OnGrowthHelper",
	1435: "CUIContext_OnContentsMap",
	1436: "CUIContext_OnUrusShop",
	1437: "CUIContext_OnDailyGift",
	1438: "CUIContext_SlidePuzzleNetworkOnPacket",
	1439: "CUIContext_OnDisguise",
	1440: "CUIContext_UNKStart",
	1451: "CUIContext_UNKEnd",
	1480: "CFuncKeyMappedMan_OnInit",
	1481: "CFuncKeyMappedMan_OnPetConsumeItemID",
	1482: "CFuncKeyMappedMan_OnPetConsumeMPItemID",
	1483: "CFuncKeyMappedMan_OnPetConsumeSkillID",
	1484: "CFuncKeyMappedMan_OnPetConsumeUNKID",
}
