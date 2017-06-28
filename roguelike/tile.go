package roguelike

type Tile interface {
	IsBlocked() bool
	IsOpaque() bool

	SetBlocked()
	SetOpaque()
}

type tile struct {
	blocked bool
	opaque  bool
}

func (t *tile) IsBlocked() bool {
	return t.blocked
}

func (t *tile) IsOpaque() bool {
	return t.opaque
}

func (t *tile) SetBlocked() {
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
