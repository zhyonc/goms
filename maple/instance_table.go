package maple

type InstanceTable struct {
	Name string
	Col  int32
	Row  int32
}

func NewInstanceTable(tableName string, col, row int32) InstanceTable {
	t := InstanceTable{
		Name: tableName,
		Col:  col,
		Row:  row,
	}
	return t
}

func (t *InstanceTable) GetValue() int32 {
	var value int32 = 0
	switch t.Name {
	case "hyper":
		if t.Row == 1 {
			value = getHyperActiveSkillSpByLv(t.Col)
		} else {
			value = getHyperPassiveSkillSpByLv(t.Col)
		}
	case "incHyperStat":
		value = getHyperStatSpByLv(t.Col)
	case "needHyperStatLv":
		value = getNeededSpForHyperStatSkill(t.Col)
	case "92000000",
		"92010000",
		"92020000",
		"92030000",
		"92040000":
		value = 100
	}
	return value
}

func getHyperActiveSkillSpByLv(level int32) int32 {
	if level == 150 || level == 170 || level == 200 {
		return 1
	}
	return 0
}

func getHyperPassiveSkillSpByLv(level int32) int32 {
	if level >= 150 && level <= 220 && level%10 == 0 {
		return 1
	}
	return 0
}

func getHyperStatSpByLv(level int32) int32 {
	return int32(3 + ((level - 140) / 10))
}

func getNeededSpForHyperStatSkill(level int32) int32 {
	switch level {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	case 4:
		return 8
	case 5:
		return 10
	case 6:
		return 15
	case 7:
		return 20
	case 8:
		return 25
	case 9:
		return 30
	case 10:
		return 35
	default:
		return 0
	}
}
