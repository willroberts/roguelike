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
			if !g.Level.IsBlocked(g.Player.X(), g.Player.Y()-1) {
				g.Player.Move(0, -1)
			}
		}
		if ev.Key == termbox.KeyArrowDown {
			if !g.Level.IsBlocked(g.Player.X(), g.Player.Y()+1) {
				g.Player.Move(0, 1)
			}
		}
		if ev.Key == termbox.KeyArrowLeft {
			if !g.Level.IsBlocked(g.Player.X()-1, g.Player.Y()) {
				g.Player.Move(-1, 0)
			}
		}
		if ev.Key == termbox.KeyArrowRight {
			if !g.Level.IsBlocked(g.Player.X()+1, g.Player.Y()) {
				g.Player.Move(1, 0)
			}
		}
	}
}
