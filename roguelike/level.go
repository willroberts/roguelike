package roguelike

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

type Level interface {
	Tiles() [][]Tile
	IsBlocked(x, y int) bool
	Render()
}

type level struct {
	w     int
	h     int
	tiles [][]Tile
}

func (l *level) Tiles() [][]Tile {
	return l.tiles
}

func (l *level) IsBlocked(x, y int) bool {
	return l.tiles[x][y].IsBlocked()
}

func (l *level) Render() {
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			t := l.tiles[x][y]
			if t.IsBlocked() {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, colorBrown)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, colorGreen)
			}
		}
	}
}

func NewLevel(w, h int) Level {
	fmt.Printf("A")
	tiles := make([][]Tile, 0)
	for x := 0; x < w; x++ {
		column := []Tile{}
		for y := 0; y < h; y++ {
			column = append(column, NewTile(false, false))
		}
		tiles = append(tiles, column)
	}
	return &level{
		w:     w,
		h:     h,
		tiles: tiles,
	}
}
