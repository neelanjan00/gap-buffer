package gapbuffer

type GapBuffer struct {
	buffer     []rune
	preGapLen  int
	postGapLen int
}

func (g *GapBuffer) gapStart() int {

	return g.preGapLen
}

func (g *GapBuffer) gapLen() int {

	return g.postGapStart() - g.preGapLen
}

func (g *GapBuffer) postGapStart() int {

	return len(g.buffer) - g.postGapLen
}

func (g *GapBuffer) SetText(s string) {

	g.buffer = []rune(s)
	g.preGapLen = 0
	g.postGapLen = len(g.buffer)
}

func (g *GapBuffer) GetText() string {

	text := append([]rune{}, g.buffer[:g.preGapLen]...)
	text = append(text, g.buffer[g.postGapStart():]...)

	return string(text)
}

func (g *GapBuffer) MoveCursorRight() {

	if g.postGapLen == 0 {
		return
	}

	g.buffer[g.preGapLen] = g.buffer[g.postGapStart()]
	g.preGapLen++
	g.postGapLen--
}

func (g *GapBuffer) MoveCursorLeft() {

	if g.preGapLen == 0 {
		return
	}

	g.buffer[g.postGapStart()-1] = g.buffer[g.preGapLen-1]
	g.preGapLen--
	g.postGapLen++
}

func (g *GapBuffer) Delete() {

	if g.postGapLen == 0 {
		return
	}

	g.postGapLen--
}

func (g *GapBuffer) Backspace() {

	if g.preGapLen == 0 {
		return
	}

	g.preGapLen--
}

func (g *GapBuffer) growGap() {

	newBuffer := make([]rune, len(g.buffer)*2)

	copy(newBuffer, g.buffer[:g.preGapLen])
	copy(newBuffer[g.postGapStart()+len(g.buffer):], g.buffer[g.postGapStart():])

	g.buffer = newBuffer
}

func (g *GapBuffer) Insert(c rune) {

	if g.gapLen() == 0 {
		g.growGap()
	}

	g.buffer[g.gapStart()] = c
	g.preGapLen++
}
