package wordwrap_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/wordwrap"
	"github.com/stretchr/testify/assert"
)

func TestNewColumnWidth_ShouldNotInstanceColumnWidth_WhenValueIsEqualToZero(t *testing.T) {
	columnWidth, err := wordwrap.NewColumnWidth(0)

	assert.Equal(t, wordwrap.NoColumnWidth, columnWidth)
	assert.EqualError(t, wordwrap.NoneZeroValueErr, err.Error())
}
