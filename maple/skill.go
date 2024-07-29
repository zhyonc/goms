package maple

// Call by CharacterData::Decode
// is_skill_need_master_level
func IsSkillNeedMasterLevel(skillID uint32) bool {
	if IsIgnoreMasterLevel(skillID) ||
		skillID/1000000 == 92 && skillID%10000 == 0 ||
		IsMakingSkillRecipe(skillID) ||
		IsCommonSkill(skillID) ||
		IsNoviceSkill(skillID) ||
		IsFieldAttackObjSkill(skillID) {
		return false
	}
	jobID := GetSkillRootFromSkill(skillID)
	return skillID != 42120024 && !IsBeastTamer(uint16(jobID)) && IsAddedSPDualAndZeroSkill(skillID) ||
		GetJobLevel(uint16(jobID)) == 4 || !IsZeroJob(uint16(jobID))
}

// is_ignore_master_level
func IsIgnoreMasterLevel(skillID uint32) bool {
	switch skillID {
	case 1120012,
		1320011,
		2121009,
		2221009,
		2321010,
		3210015,
		4110012,
		4210012,
		4340010,
		4340012,
		5120011,
		5120012,
		5220012,
		5220014,
		5320007,
		5321004,
		5321006,
		21120011,
		21120014,
		21120020,
		21120021,
		21121008,
		22171069,
		23120011,
		23120013,
		23121008,
		33120010,
		35120014,
		51120000,
		80001913:
		return true
	}
	return false
}

// is_making_skill_recipe
func IsMakingSkillRecipe(recipeID uint32) bool {
	if recipeID/1000000 != 92 || recipeID%10000 == 1 {
		v1 := 10000 * (recipeID / 10000)
		if v1/1000000 == 92 && (v1%10000 == 0) {
			return true
		}
	}
	return false
}

// is_common_skill
func IsCommonSkill(skillID uint32) bool {
	v1 := skillID / 10000
	if skillID/10000 == 8000 {
		v1 = skillID / 100
	}
	return v1 >= 800000 && v1 <= 800099
}

// is_novice_skill
func IsNoviceSkill(skillID uint32) bool {
	v1 := skillID / 10000
	if skillID/10000 == 8000 {
		v1 = skillID / 100
	}
	return IsBeginnerJob(uint16(v1))
}

// is_field_attack_obj_skill
func IsFieldAttackObjSkill(skillID uint32) bool {
	if skillID <= 0 {
		return false
	}
	v1 := skillID / 10000
	if skillID/10000 == 8000 {
		v1 = skillID / 100
	}
	return v1 == 9500
}

// get_skill_root_from_skill
func GetSkillRootFromSkill(skillID uint32) uint32 {
	result := skillID / 10000
	if skillID/10000 == 8000 {
		return skillID / 100
	}
	return result
}

// is_added_sp_dual_and_zero_skill
func IsAddedSPDualAndZeroSkill(skillID uint32) bool {
	switch skillID {
	case 4311003,
		4321006,
		4330009,
		4331002,
		4340007,
		4341004,
		101000101,
		101100101,
		101100201,
		101110102,
		101110200,
		101110203,
		101120104,
		101120204:
		return true
	}
	return false
}
