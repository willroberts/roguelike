package main

import (
	"log"

	termbox "github.com/nsf/termbox-go"
	"github.com/willroberts/roguelike/roguelike"
)

func main() {
	// Create and initialize our game.
	g := &roguelike.Game{}
	if err := g.Setup(); err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	// Run the game loop until the game sets g.Running to false.
	for g.Running {
		select {
		case ev := <-g.EventQueue:
			g.HandleInput(ev)
		default:
			g.Main()
		}
	}
}
