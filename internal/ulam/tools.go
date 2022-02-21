package ulam

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

// Convert 0-65535 color to 0-1 color
func col(c uint32) float64 {
	return float64(c) / 65535
}

func getCenter(da *gtk.DrawingArea) (w int, h int) {
	h = da.GetAllocatedHeight() / 2
	w = da.GetAllocatedWidth() / 2
	return
}

func getMovement(d direction) (x int, y int) {
	switch d {
	case up:
		return 0, -stepY
	case left:
		return -stepX, 0
	case down:
		return 0, stepY
	case right:
		return stepX, 0
	default:
		panic("invalid direction")
	}
}

func isPrime(p int) bool {
	if p == 0 || p == 1 {
		fmt.Println(p, false)
		return false
	}

	for i := 2; i <= p/2; i++ {
		if p%i == 0 {
			fmt.Println(p, false)
			return false
		}
	}
	fmt.Println(p, true)
	return true
}
