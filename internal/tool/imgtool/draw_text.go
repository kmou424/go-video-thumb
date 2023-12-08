package imgtool

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"math"
)

var DefaultFontDPI float64 = 72

type TextOption struct {
	X        int
	Y        int
	FontData []byte
	Size     float64
	String   string
}

func DrawText(rgba *image.RGBA, option *TextOption) error {
	f, err := truetype.Parse(option.FontData)
	if err != nil {
		return err
	}
	face := truetype.NewFace(f, &truetype.Options{
		DPI:     DefaultFontDPI,
		Size:    option.Size,
		Hinting: font.HintingFull,
	})
	drawer := font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(image.Black),
		Face: face,
		Dot:  fixed.P(option.X, option.Y),
	}
	drawer.DrawString(option.String)

	return nil
}

func DrawTextOutline(rgba *image.RGBA, option *TextOption) error {
	f, err := truetype.Parse(option.FontData)
	if err != nil {
		return err
	}
	face := truetype.NewFace(f, &truetype.Options{
		DPI:     DefaultFontDPI,
		Size:    option.Size,
		Hinting: font.HintingFull,
	})

	drawer := font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(image.White),
		Face: face,
		Dot:  fixed.P(option.X, option.Y),
	}

	borderWidth := int(math.Ceil(option.Size / 10.0))
	for dy := -borderWidth; dy <= borderWidth; dy++ {
		for dx := -borderWidth; dx <= borderWidth; dx++ {
			if dx != 0 || dy != 0 {
				drawer.Src = image.NewUniform(image.Black)
				drawer.Dot = fixed.P(option.X+dx, option.Y+dy)
				drawer.DrawString(option.String)
			}
		}
	}

	drawer.Src = image.NewUniform(image.White)
	drawer.Dot = fixed.P(option.X, option.Y)
	drawer.DrawString(option.String)

	return nil
}
