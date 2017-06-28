package roguelike

import (
	termbox "github.com/nsf/termbox-go"
)

// HandleInput processes termbox events and changes game state accordingly.
func (g *Game) HandleInput(ev termbox.Event) {
	if ev.Type == termbox.EventKey {
		if ev.Key == termbox.KeyEsc {
			g.Running = false
		}
		if ev.Key == termbox.KeyArrowUp {
			g.Player.Move(0, -1)
		}
		if ev.Key == termbox.KeyArrowDown {
			g.Player.Move(0, 1)
		}
		if ev.Key == termbox.KeyArrowLeft {
			g.Player.Move(-1, 0)
		}
		if ev.Key == termbox.KeyArrowRight {
			g.Player.Move(1, 0)
		}
	}
}
