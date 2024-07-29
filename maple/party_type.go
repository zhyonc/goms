package maple

type PartyType uint8

const (
	PartyReq_LoadParty PartyType = iota
	PartyReq_CreateNewParty
	PartyReq_WithdrawParty
	PartyReq_JoinParty
	PartyReq_InviteParty
	PartyReq_InviteIntrusion // member -> party request
	PartyReq_KickParty       // party -> member request
	PartyReq_ChangePartyBoss
	PartyReq_ApplyParty
	PartyReq_SetAppliable
	PartyReq_ClearIntrusion
	PartyReq_CreateNewParty_Group
	PartyReq_JoinParty_Group
	PartyReq_PartySetting
	PartyReq_LoadStarPlanetPoint
	_
	PartyRes_LoadParty_Done
	PartyRes_CreateNewParty_Done
	PartyRes_CreateNewParty_AlreayJoined
	PartyRes_CreateNewParty_Beginner
	PartyRes_CreateNewParty_Unknown
	PartyRes_CreateNewParty_byNonBoss

	PartyRes_WithdrawParty_Done
	PartyRes_WithdrawParty_NotJoined
	PartyRes_WithdrawParty_Unknown

	PartyRes_JoinParty_Done
	PartyRes_JoinParty_Done2

	PartyRes_JoinParty_AlreadyJoined
	PartyRes_JoinParty_AlreadyFull
	PartyRes_JoinParty_OverDesiredSize
	PartyRes_JoinParty_UnknownUser
	PartyRes_JoinParty_Unknown

	PartyRes_JoinIntrusion_Done
	PartyRes_JoinIntrusion_UnknownParty

	PartyRes_InviteParty_Sent
	PartyRes_InviteParty_BlockedUser
	PartyRes_InviteParty_AlreadyInvited
	PartyRes_InviteParty_AlreadyInvitedByInviter
	PartyRes_InviteParty_Rejected
	PartyRes_InviteParty_Accepted

	PartyRes_InviteIntrusion_Sent
	PartyRes_InviteIntrusion_BlockedUser
	PartyRes_InviteIntrusion_AlreadyInvited
	PartyRes_InviteIntrusion_AlreadyInvitedByInviter
	PartyRes_InviteIntrusion_Rejected
	PartyRes_InviteIntrusion_Accepted

	PartyRes_KickParty_Done
	PartyRes_KickParty_FieldLimit
	PartyRes_KickParty_Unknown
	PartyRes_KickParty_Unavailable

	PartyRes_ChangePartyBoss_Done
	PartyRes_ChangePartyBoss_NotSameField
	PartyRes_ChangePartyBoss_NoMemberInSameField
	PartyRes_ChangePartyBoss_NotSameChannel
	PartyRes_ChangePartyBoss_Unknown

	PartyRes_AdminCannotCreate
	PartyRes_AdminCannotInvite

	PartyRes_InAnotherWorld
	PartyRes_InAnotherChanelBlockedUser

	PartyRes_UserMigration
	PartyRes_ChangeLevelOrJob
	PartyRes_UpdateShutdownStatus
	PartyRes_UNK62
	PartyRes_SetAppliable
	PartyRes_SetAppliableFailed
	PartyRes_SuccessToSelectPQReward
	PartyRes_FailToSelectPQReward
	PartyRes_ReceivePQReward
	PartyRes_FailToRequestPQReward
	PartyRes_CanNotInThisField

	PartyRes_ApplyParty_Sent
	PartyRes_ApplyParty_UnknownParty
	PartyRes_ApplyParty_BlockedUser
	PartyRes_ApplyParty_AlreadyApplied
	PartyRes_ApplyParty_AlreadyAppliedByApplier
	PartyRes_ApplyParty_AlreadyFull
	PartyRes_ApplyParty_Rejected
	PartyRes_ApplyParty_Accepted

	PartyRes_FoundPossibleMember
	PartyRes_FoundPossibleParty

	PartyRes_PartySettingDone
	PartyRes_Load_StarGrade_Result
	PartyRes_Load_StarGrade_Result2
	PartyRes_Member_Rename
	PartyRes_UNK84
	PartyRes_UNK85

	PartyInfo_TownPortalChanged
	PartyInfo_OpenGate

	// below enums unchecked
	ExpeditionReq_Load
	ExpeditionReq_CreateNew
	ExpeditionReq_Invite
	ExpeditionReq_ResponseInvite
	ExpeditionReq_Withdraw
	ExpeditionReq_Kick
	ExpeditionReq_ChangeMaster
	ExpeditionReq_ChangePartyBoss
	ExpeditionReq_RelocateMember

	ExpeditionNoti_Load_Done
	ExpeditionNoti_Load_Fail
	ExpeditionNoti_CreateNew_Done
	ExpeditionNoti_Join_Done
	ExpeditionNoti_You_Joined
	ExpeditionNoti_You_Joined2
	ExpeditionNoti_Join_Fail
	ExpeditionNoti_Withdraw_Done
	ExpeditionNoti_You_Withdrew
	ExpeditionNoti_Kick_Done
	ExpeditionNoti_You_Kicked
	ExpeditionNoti_Removed
	ExpeditionNoti_MasterChanged
	ExpeditionNoti_Modified
	ExpeditionNoti_Modified2
	ExpeditionNoti_Invite
	ExpeditionNoti_ResponseInvite
	ExpeditionNoti_Create_Fail_By_Over_Weekly_Counter
	ExpeditionNoti_Invite_Fail_By_Over_Weekly_Counter
	ExpeditionNoti_Apply_Fail_By_Over_Weekly_Counter
	ExpeditionNoti_Invite_Fail_By_Blocked_Behavior

	AdverNoti_LoadDone
	AdverNoti_Change
	AdverNoti_Remove
	AdverNoti_GetAll
	AdverNoti_Apply
	AdverNoti_ResultApply
	AdverNoti_AddFail
	AdverReq_Add
	AdverReq_Remove
	AdverReq_GetAll
	AdverReq_RemoveUserFromNotiList
	AdverReq_Apply
	AdverReq_ResultApply
)