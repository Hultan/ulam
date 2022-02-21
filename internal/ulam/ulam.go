package ulam

import (
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Ulam struct {
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea

	tickerQuit chan struct{}
	ticker     *time.Ticker
	speed      time.Duration
	isActive   bool

	primes []int
}

func NewUlam(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *Ulam {
	t := &Ulam{window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)
	// t.window.Connect("button-press-event", t.onKeyPressed)
	// t.window.Connect("event", t.onKeyPressed)

	return t
}

func (p *Ulam) StartGame() {
	p.window.Maximize()
	p.drawingArea.Connect("draw", p.onDraw)
	p.speed = 500
	p.ticker = time.NewTicker(p.speed * time.Millisecond)
	p.tickerQuit = make(chan struct{})

	p.primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 27}

	go p.mainLoop()
}

func (p *Ulam) mainLoop() {
	for {
		select {
		case <-p.ticker.C:
			p.drawingArea.QueueDraw()
		case <-p.tickerQuit:
			p.isActive = false
			p.ticker.Stop()
			return
		}
	}
}

// onKeyPressed : The onKeyPressed signal handler
func (p *Ulam) onKeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	switch key.KeyVal() {
	case 113: // Button "Q" => Quit game
		p.quit()
		p.window.Close() // Close window
	}
	p.drawingArea.QueueDraw()
}

func (p *Ulam) quit() {
	if p.isActive {
		p.isActive = false
		close(p.tickerQuit) // Stop ticker
	}
}
