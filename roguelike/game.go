package roguelike

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

const (
	fTime     time.Duration     = 1000 / 30 * time.Millisecond // 30 FPS.
	bgColor   termbox.Attribute = termbox.ColorBlack
	wallColor termbox.Attribute = termbox.ColorGreen
)

// Game is our container for all game data. We bind methods to it in order to
// avoid creating global variables and global state.
type Game struct {
	Running    bool
	EventQueue chan termbox.Event
	Level      Level
	Player     Entity
	Entities   []Entity
}

// Setup performs one-time tasks at startup in order to initialize the game.
func (g *Game) Setup() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	termbox.SetOutputMode(termbox.Output256)

	g.EventQueue = make(chan termbox.Event)
	go func(g *Game) {
		for {
			g.EventQueue <- termbox.PollEvent()
		}
	}(g)

	g.Level = NewLevel(80, 45)
	g.Level.Tiles()[30][22].SetBlocking()
	g.Level.Tiles()[30][22].SetOpaque()
	g.Level.Tiles()[31][22].SetBlocking()
	g.Level.Tiles()[31][22].SetOpaque()
	g.Level.Tiles()[32][22].SetBlocking()
	g.Level.Tiles()[32][22].SetOpaque()

	g.Player = NewEntity(1, 1, charRogue, termbox.ColorYellow)
	g.Entities = []Entity{
		NewEntity(16, 16, charLight, termbox.ColorWhite),
	}
	g.Running = true

	return nil
}

// Main is the game's core loop. It is expected to be run from a main package
// until g.Running is set to false by the game.
func (g *Game) Main() {
	termbox.Clear(termbox.ColorDefault, bgColor)

	// Draw player and other entities.
	g.Player.Render()
	for _, e := range g.Entities {
		e.Render()
	}

	termbox.Flush()
	time.Sleep(fTime)
}
