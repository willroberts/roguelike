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
			g.PlayerY--
		}
		if ev.Key == termbox.KeyArrowDown {
			g.PlayerY++
		}
		if ev.Key == termbox.KeyArrowLeft {
			g.PlayerX--
		}
		if ev.Key == termbox.KeyArrowRight {
			g.PlayerX++
		}
	}
}
