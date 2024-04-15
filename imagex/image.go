package imagex

import (
	"github.com/nfnt/resize"
	"image"
)

func Resize(width, height uint, img image.Image) image.Image {
	return resize.Resize(width, height, img, resize.Bicubic)
}

func Thumbnail(maxWidth, maxHeight uint, img image.Image) image.Image {
	return resize.Thumbnail(maxWidth, maxHeight, img, resize.Bicubic)
}
