package wordwrap_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/wordwrap"
	"github.com/stretchr/testify/assert"
)

func TestWordWrap_ShouldReturnCurrentString_WhenTextLengthIsLowerThanWidthLimit(t *testing.T) {
	columnWidth, err := wordwrap.NewColumnWidth(10)

	actual := wordwrap.Wrap("Chacho", columnWidth)

	assert.Equal(t, "Chacho", actual)
	assert.Empty(t, err)
}

func TestWordWrap_ShouldReturnWrappedString_WhenTextLengthIsGreaterThanWidthLimit(t *testing.T) {
	columnWidth, err := wordwrap.NewColumnWidth(10)

	actual := wordwrap.Wrap("0123456789012345", columnWidth)

	assert.Equal(t, "0123456789\n012345", actual)
	assert.Empty(t, err)
}

func TestWordWrap_ShouldReturnWrappedStringWithMultiple_WhenTextLengthIsGreaterThanWidthLimit(t *testing.T) {
	columnWidth, err := wordwrap.NewColumnWidth(10)

	actual := wordwrap.Wrap("0123456789012345678901234", columnWidth)

	assert.Equal(t, "0123456789\n0123456789\n01234", actual)
	assert.Empty(t, err)
}

func TestWordWrap_ShouldReturnError_WhenColumnWidthIsLowerOrEqualThanZero(t *testing.T) {
	columnWidth, err := wordwrap.NewColumnWidth(10)

	actual := wordwrap.Wrap("0123456789012345678901234", columnWidth)

	assert.Equal(t, "0123456789\n0123456789\n01234", actual)
	assert.Empty(t, err)
}
