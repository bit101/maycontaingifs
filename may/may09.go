package may

import (
	"fmt"
	"math"
	"os/exec"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
	"github.com/bit101/bitlibgo/random"
)

// May09 generates a gif
func May09() {
	width := 400.0
	height := 400.0
	filename := "out/may09.gif"

	type Point3D struct {
		x float64
		y float64
		z float64
	}

	var points []*Point3D
	for i := 0; i < 100; i++ {
		angle := random.FloatRange(0, math.Pi)
		radius := random.FloatRange(-250, 250)
		p := &Point3D{
			math.Cos(angle) * radius,
			random.FloatRange(-200, 200),
			math.Sin(angle) * radius,
		}
		points = append(points, p)
	}

	fl := 300.0

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 180
	animation.Render(func(surface *bitlibgo.BitSurface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		surface.Save()
		surface.Translate(width/2, height/2)
		angle := percent * math.Pi * 2

		for _, p := range points {
			xx := math.Cos(angle)*p.x - math.Sin(angle)*p.z
			zz := math.Cos(angle)*p.z + math.Sin(angle)*p.x
			scale := fl / (fl + zz + 200)
			x := scale * xx
			y := scale * p.y
			r := scale * 30

			surface.FillCircle(x, y, r)
		}
		surface.Restore()
	})
	cmd := exec.Command("cp", filename, "out/latest.gif")
	cmd.Run()
}
