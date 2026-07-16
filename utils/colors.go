package utils

import (
	"image"
	"image/color/palette"
	"image/draw"
)

func GetPixelColorIndices(img image.Image) []uint8 {
	// using Plan9 means index values will always be between 0-255
	paletted := image.NewPaletted(img.Bounds(), palette.Plan9)

	// this could be a bottleneck for large images
	draw.Draw(paletted, paletted.Rect, img, img.Bounds().Min, draw.Src)

	return paletted.Pix
}

func GetPadFromImage(path string) []uint8 {
	image := ReadJpegPhoto(path)
	pads := GetPixelColorIndices(image)

	return pads
}
