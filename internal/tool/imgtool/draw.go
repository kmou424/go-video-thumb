package imgtool

import (
	"errors"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
)

func DrawWhiteImage(height int, width int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(),
		&image.Uniform{
			C: color.White,
		}, image.Point{}, draw.Src)

	return img
}

func DrawImageResize(background *image.RGBA, imageToDraw image.Image, height, width, x, y int) {
	resizedImageToDraw := imaging.Resize(imageToDraw, width, height, imaging.Lanczos)
	rect := image.Rect(x, y, x+width, y+height)
	draw.Draw(background, rect, resizedImageToDraw, image.Point{}, draw.Src)
}

func ConcatImages(images ...image.Image) (image.Image, error) {
	if len(images) == 0 {
		return nil, errors.New("no images provided")
	}

	firstImage := images[0]
	width := firstImage.Bounds().Size().X
	height := 0

	for _, img := range images {
		height += img.Bounds().Size().Y
	}

	result := image.NewRGBA(image.Rect(0, 0, width, height))
	yOffset := 0

	for _, img := range images {
		imgHeight := img.Bounds().Size().Y

		draw.Draw(result, image.Rect(0, yOffset, width, yOffset+imgHeight), img, image.Point{}, draw.Src)

		yOffset += imgHeight
	}

	return result, nil
}
