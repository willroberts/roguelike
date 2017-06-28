package roguelike

import termbox "github.com/nsf/termbox-go"

const (
	charRogue rune = '@'
	charLight rune = '☺'
	charDark  rune = '☻'

	colorPink termbox.Attribute = 0xb8
)

var (
	blockRunes = []rune{' ', '░', '▒', '▓', '█'}
	neatRunes  = []rune{'※', '⁜', '⁕', '⁂'}
	// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Block_Elements
	// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Geometric_Shapes
	// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Miscellaneous_Symbols
	// https://en.wikipedia.org/wiki/Emoji#Unicode_blocks
)
