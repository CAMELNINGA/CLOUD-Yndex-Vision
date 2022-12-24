package croppimage

import (
	"image"

	"github.com/oliamb/cutter"
)

func CropFase(img image.Image) (image.Image, error) {
	return cutter.Crop(img, cutter.Config{
		Width:  250,
		Height: 500,
		Anchor: image.Point{100, 100},
		Mode:   cutter.TopLeft, // optional, default value
	})
}
