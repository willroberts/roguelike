package main

import (
	"log"
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

const (
	fTime time.Duration = 1000 / 30 * time.Millisecond // 30 FPS.
)

var runes = []rune{' ', '░', '▒', '▓', '█'}

func DrawScreen() {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r := runes[rand.Intn(len(runes))]
			attr := rand.Int()%8 + 1
			termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.Attribute(attr))
		}
	}
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatalln("error initializing termbox:", err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	DrawScreen()

	var running bool = true
	for running {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				running = false
			}
		default:
			DrawScreen()
			time.Sleep(fTime)
		}
	}
}
