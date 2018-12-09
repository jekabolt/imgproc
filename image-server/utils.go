package server

import (
	"image"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

func cropImage(input image.Image, dimension int) (image.Image, error) {
	// dx horizontal
	// dy vertical
	horizontal := input.Bounds().Dx()
	vertical := input.Bounds().Dy()

	// Square
	if horizontal == vertical {
		// compress image and return with needed dimension
		resized := resize.Resize(uint(600), 0, input, resize.Bilinear)
		return resize.Resize(uint(dimension), 0, resized, resize.Bilinear), nil
	}

	// Portrait
	if vertical > horizontal {
		// compress image
		resized := resize.Resize(uint(600), 0, input, resize.Bilinear)
		img := resize.Resize(uint(dimension), 0, resized, resize.Bilinear)

		croppedImg, err := cutter.Crop(img, cutter.Config{
			Width:  dimension,
			Height: dimension,
			Anchor: image.Point{0, 0},
			Mode:   cutter.Centered,
		})
		return croppedImg, err

	}

	// Landscape
	if horizontal > vertical {
		resized := resize.Resize(0, uint(600), input, resize.Bilinear)
		img := resize.Resize(0, uint(dimension), resized, resize.Bilinear)

		croppedImg, err := cutter.Crop(img, cutter.Config{
			Width:  dimension,
			Height: dimension,
			Anchor: image.Point{0, 0},
			Mode:   cutter.Centered,
		})
		return croppedImg, err

	}
	return input, nil
}
