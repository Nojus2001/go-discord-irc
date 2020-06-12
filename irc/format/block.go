package ircf

import "fmt"

// From https://www.npmjs.com/package/irc-formatting 1.0.0-rc3

type Block struct {
	Bold, Italic, Underline, Reverse bool
	Color, Highlight                 int
	Text                             string
}

var Empty = NewBlock("")

func NewBlock(text string, fields ...rune) (this Block) {
	this.Text = text
	this.Color = -1
	this.Highlight = -1

	for _, code := range fields {
		this.SetField(code, true)
	}

	return
}

func NewColorBlock(text string, fg int, bg int, fields ...rune) Block {
	b := NewBlock(text, fields...)
	b.Color = fg
	b.Highlight = bg
	return b
}

func (this Block) Equals(other Block) bool {
	return this.Bold == other.Bold &&
		this.Italic == other.Italic &&
		this.Underline == other.Underline &&
		this.Reverse == other.Reverse &&
		this.Color == other.Color &&
		this.Highlight == other.Highlight
}

func (this Block) IsPlain() bool {
	return (!this.Bold && !this.Italic && !this.Underline && !this.Reverse &&
		this.Color == -1 && this.Highlight == -1)
}

func (this Block) HasSameColor(other Block, reversed bool) bool {
	if this.Reverse && reversed {
		return ((this.Color == other.Highlight || other.Highlight == -1) && this.Highlight == other.Color)
	}
	return (this.Color == other.Color && this.Highlight == other.Highlight)
}

func (this Block) GetColorString() string {
	var str = ""

	if this.Color != -1 {

		str = fmt.Sprintf("%02d", this.Color)
	}

	if this.Highlight != -1 {
		str += "," + fmt.Sprintf("%02d", this.Highlight)
	}

	return str
}

func (this *Block) codeToField(code rune) (field *bool) {
	if code == CharBold {
		field = &this.Bold
	} else if code == CharItalics {
		field = &this.Italic
	} else if code == CharUnderline {
		field = &this.Underline
	} else if code == CharReverseColor {
		field = &this.Reverse
	}
	return field
}

func (this *Block) SetField(code rune, val bool) {
	if field := this.codeToField(code); field != nil {
		*field = val
		return
	}
	panic(fmt.Sprintf(`Unknown code \x%x`, code))
}

func (this Block) GetField(code rune) bool {
	if field := this.codeToField(code); field != nil {
		return *field
	}
	panic(fmt.Sprintf(`Unknown code \x%x`, code))
}
