package roguelike

import termbox "github.com/nsf/termbox-go"

const (
	charRogue rune = '@'
	charLight rune = '☺'
	charDark  rune = '☻'

	colorBlack     termbox.Attribute = 0x11
	colorBrown     termbox.Attribute = 0x60
	colorGold      termbox.Attribute = 0x90
	colorGreen     termbox.Attribute = 0x17
	colorIndigo    termbox.Attribute = 0x40
	colorLtOrange  termbox.Attribute = 0xe0
	colorLtPurple  termbox.Attribute = 0xb8
	colorLtSeafoam termbox.Attribute = 0xa0
	colorMagenta   termbox.Attribute = 0x80
	colorPink      termbox.Attribute = 0xd0
	colorSalmon    termbox.Attribute = 0xb0
	colorSeafoam   termbox.Attribute = 0x50
	colorSkyBlue   termbox.Attribute = 0x70
	colorTeal      termbox.Attribute = 0x20
	colorWhite     termbox.Attribute = 0xe8
	colorYellow    termbox.Attribute = 0xc0

	bgColor termbox.Attribute = colorGreen
)

var (
	blockRunes = []rune{' ', '░', '▒', '▓', '█'}
	neatRunes  = []rune{'※', '⁜', '⁕', '⁂'}
	// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Block_Elements
	// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Geometric_Shapes
	// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Miscellaneous_Symbols
	// https://en.wikipedia.org/wiki/Emoji#Unicode_blocks
)
