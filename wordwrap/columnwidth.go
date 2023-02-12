package wordwrap

import (
	"errors"
	"fmt"
)

var (
	NoColumnWidth   = ColumnWidth{}
	InvalidValueErr = func(value int) error { return errors.New(fmt.Sprintf("invalid value %d", value)) }
)

type ColumnWidth struct {
	value int
}

func (c ColumnWidth) Value() int {
	return c.value
}

func NewColumnWidth(value int) (ColumnWidth, error) {
	if value <= 0 {
		return NoColumnWidth, InvalidValueErr(value)
	}
	return ColumnWidth{value}, nil
}
