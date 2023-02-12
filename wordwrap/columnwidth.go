package wordwrap

import (
	"errors"
)

var (
	NoColumnWidth    = ColumnWidth{}
	NoneZeroValueErr = errors.New("value cannot be zero")
)

type ColumnWidth struct {
	value uint
}

func (c ColumnWidth) Value() uint {
	return c.value
}

func NewColumnWidth(value uint) (ColumnWidth, error) {
	if value == 0 {
		return NoColumnWidth, NoneZeroValueErr
	}
	return ColumnWidth{value}, nil
}
