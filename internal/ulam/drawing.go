package ulam

import (
	"image/color"
	"math"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

type direction int

const (
	right direction = iota
	up
	left
	down
)

const (
	stepX = 20
	stepY = 20
)

// onDraw : The onDraw signal handler
func (p *Ulam) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	p.drawBackground(da, ctx)
	p.drawPrimes(da, ctx)
}

//
// HELPER FUNCTIONS
//

// drawBackground : Draws the background
func (p *Ulam) drawBackground(da *gtk.DrawingArea, ctx *cairo.Context) {
	width := float64(da.GetAllocatedWidth())
	height := float64(da.GetAllocatedHeight())
	p.setColor(ctx, color.White)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

func (p *Ulam) drawPrimes(da *gtk.DrawingArea, ctx *cairo.Context) {
	p.setColor(ctx, color.Black)
	cw, ch := getCenter(da)
	var x, y = cw, ch
	var dir = right
	var steps = 0
	var stepsInDirection = 1
	var count = 0

	prevX, prevY := x, y

	for i := 0; i < 2000; i++ {
		dx, dy := getMovement(dir)
		x += dx
		y += dy
		if isPrime(i + 2) {
			ctx.Arc(float64(x), float64(y), 10, 0, 2*math.Pi)
			ctx.Fill()
		}
		ctx.MoveTo(float64(prevX), float64(prevY))
		ctx.LineTo(float64(x), float64(y))
		ctx.Stroke()
		prevX, prevY = x, y

		steps += 1
		if steps == stepsInDirection {
			count += 1
			steps = 0
			dir = (dir + 1) % 4
			if count%2 == 0 {
				count = 0
				stepsInDirection += 1
			}
		}
	}
}

func (p *Ulam) setColor(ctx *cairo.Context, c color.Color) {
	r, g, b, a := c.RGBA()
	ctx.SetSourceRGBA(col(r), col(g), col(b), col(a))
}
