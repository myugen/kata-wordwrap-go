package wordwrap

const lineBreak = "\n"

func Wrap(text string, columnWidth ColumnWidth) string {
	width := columnWidth.Value()
	if uint(len(text)) <= width {
		return text
	}
	wrapped := text[0:width]
	remainingText := text[width:]
	return wrapped + lineBreak + Wrap(remainingText, columnWidth)
}
