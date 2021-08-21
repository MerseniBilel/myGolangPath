package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black} // create a color palette

const (
	whiteIndex = 0 //index of the color white
	blackIndex = 1 //index of the black color
)

func main() {
	animateGif(os.Stdout)
}

func animateGif(out io.Writer) {
	const (
		cycles = 5     // number of complete x oscillator revolutions
		res    = 0.001 // angular resolution
		size   = 100   // image canvas cover
		nframs = 64    // number of frams
		delay  = 8     // delay between frams  in 10ms units
	)

	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframs}
	phase := 0.0 // phase deffrence
	for i := 0; i < nframs; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
