package roguelike

type Tile interface {
	IsBlocking() bool
	IsOpaque() bool

	SetBlocking()
	SetOpaque()
}

type tile struct {
	blocked bool
	opaque  bool
}

func (t *tile) IsBlocking() bool {
	return t.blocked
}

func (t *tile) IsOpaque() bool {
	return t.opaque
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
