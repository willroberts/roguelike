package roguelike

import termbox "github.com/nsf/termbox-go"

type Entity interface {
	X() int
	Y() int
	Icon() rune
	Color() termbox.Attribute

	Move(int, int)
	Render()
	Clear()
}

type entity struct {
	x     int
	y     int
	icon  rune
	color termbox.Attribute
}

func (e *entity) X() int {
	return e.x
}

func (e *entity) Y() int {
	return e.y
}

func (e *entity) Icon() rune {
	return e.icon
}

func (e *entity) Color() termbox.Attribute {
	return e.color
}

func (e *entity) Move(dx, dy int) {
	e.x += dx
	e.y += dy
}

func (e *entity) Render() {
	termbox.SetCell(e.x, e.y, e.icon, e.color, bgColor)
}

func (e *entity) Clear() {
	termbox.SetCell(e.x, e.y, ' ', termbox.ColorDefault, bgColor)
}

func NewEntity(x, y int, icon rune, color termbox.Attribute) Entity {
	return &entity{
		x:     x,
		y:     y,
		icon:  icon,
		color: color,
	}
}
