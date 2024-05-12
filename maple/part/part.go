package part

import (
	"reflect"
)

func ApplyEquipPart(inventory any) map[uint8]uint32 {
	equips := make(map[uint8]uint32)
	v := reflect.ValueOf(inventory)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		partID := EquipPartMap[fieldName]
		itemID := v.Field(i).Uint()
		if partID == 0 || itemID == 0 {
			continue
		}
		equips[partID] = uint32(itemID)
	}
	return equips
}
