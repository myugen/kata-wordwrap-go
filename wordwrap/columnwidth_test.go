package wordwrap_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/wordwrap"
	"github.com/stretchr/testify/assert"
)

func TestNewColumnWidth_ShouldNotInstanceColumnWidth_WhenValueIsLessOrEqualThanZero(t *testing.T) {
	columnWidth, err := wordwrap.NewColumnWidth(-1)

	assert.Equal(t, wordwrap.NoColumnWidth, columnWidth)
	assert.EqualError(t, wordwrap.InvalidValueErr(-1), err.Error())
}
