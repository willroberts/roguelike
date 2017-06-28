package roguelike

type Level interface {
	Tiles() [][]Tile
}

type level struct {
	w     int
	h     int
	tiles [][]Tile
}

func (l *level) Tiles() [][]Tile {
	return l.tiles
}

func NewLevel(w, h int) Level {
	tiles := make([][]Tile, 0)
	for i := 0; i < w; i++ {
		column := []Tile{}
		for j := 0; j < h; j++ {
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
