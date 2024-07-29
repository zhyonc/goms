package maple

const (
	EmptyPortalID int32 = 999999999
)

// is_banban_base_field
func IsBanBanBaseField(id uint32) bool {
	return id/10 == 10520011 || id/10 == 10520051 || id == 105200519
}
