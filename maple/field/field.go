package field

import "math"

type Field struct {
	ID     int
	Top    uint32
	Bottom uint32
	Left   uint32
	Right  uint32
}

func NewField(fieldID int) *Field {
	field := &Field{
		ID: fieldID,
	}
	return field
}

func (f *Field) GetWidth() uint32 {
	return uint32(math.Abs(float64(f.Right) - float64(f.Left)))
}

func (f *Field) GetHeight() uint32 {
	return uint32(math.Abs(float64(f.Top) - float64(f.Bottom)))
}
