package nx

import (
	"log/slog"
	"reflect"
	"strings"

	"github.com/Hucaru/gonx"
)

func MappingField(entity any, node *gonx.Node, rootNodes []gonx.Node, textLookup []string) {
	v := reflect.ValueOf(entity)
	if v.Kind() != reflect.Ptr {
		slog.Error("Failed to mapping field", "err", "Entity is not a pointer")
		return
	}
	for i := uint32(0); i < uint32(node.ChildCount); i++ {
		option := rootNodes[node.ChildID+i]
		optionName := textLookup[option.NameID]
		filedName := ToUpperFirstChar(optionName)
		field := v.Elem().FieldByName(filedName)
		if !field.IsValid() {
			continue
		}
		switch field.Type().Kind() {
		case reflect.Bool:
			field.SetBool(option.Data[0] != 0)
		case reflect.String:
			field.SetString(textLookup[gonx.DataToUint32(option.Data)])
		case reflect.Int8:
			field.SetInt(int64(option.Data[0]))
		case reflect.Int16:
			field.SetInt(int64(gonx.DataToInt16(option.Data)))
		case reflect.Int32:
			field.SetInt(int64(gonx.DataToInt32(option.Data)))
		case reflect.Int64:
			field.SetInt(gonx.DataToInt64(option.Data))
		case reflect.Float64:
			field.SetFloat(gonx.DataToFloat64(option.Data))
		default:
			slog.Warn("Unknown field type", "optionName", optionName, "optionData", option.Data)
		}
	}
}

func ToUpperFirstChar(s string) string {
	return (strings.ToUpper(s[:1]) + s[1:])
}
