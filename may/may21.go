package may

import (
	"fmt"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
	"github.com/bit101/blg/noise"
)

// May21 generates a gif
func May21() {
	width := 400.0
	height := 400.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)

		scale := 0.01
		res := 1.0

		for x := 0.0; x < width; x += res {
			for y := 0.0; y < width; y += res {
				v := blmath.Map(noise.Perlin(x*scale, y*scale, blmath.LerpSin(percent, 0, 2)), -1, 1, 0, 1)
				surface.SetSourceRGB(v, v, v)
				surface.FillRectangle(x, y, res, res)
			}
		}

	})
}
