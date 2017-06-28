package roguelike

type Tile interface {
	SetBlocking()
	SetOpaque()
}

type tile struct {
	blocked bool
	opaque  bool
}

func (t *tile) SetBlocking() {
	t.blocked = true
}

func (t *tile) SetOpaque() {
	t.opaque = true
}

func NewTile(blocked, opaque bool) Tile {
	return &tile{
		blocked: blocked,
		opaque:  opaque,
	}
}
