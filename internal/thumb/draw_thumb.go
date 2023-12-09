package thumb

import (
	"github.com/kmou424/go-video-thumb/internal/font"
	"github.com/kmou424/go-video-thumb/internal/global"
	"github.com/kmou424/go-video-thumb/internal/tool"
	"github.com/kmou424/go-video-thumb/internal/tool/imgtool"
	"image"
	"image/draw"
)

func drawThumbnailImage(frames []*imgtool.Frame) *image.RGBA {
	backgroundImg := imgtool.DrawWhiteImage(defaultFrameHeight*global.ImageRows, defaultFrameWidth*global.ImageColumns)

	for row := 0; row < global.ImageRows; row++ {
		heightOffset := defaultFrameHeight * row
		for column := 0; column < global.ImageColumns; column++ {
			frame := frames[row*global.ImageColumns+column]
			tool.Logger.Infof(
				"drawing frame (#%d)[%s] at (%d, %d)",
				frame.Count,
				tool.FormatDuration(frame.TimePointDuration),
				column*defaultFrameWidth,
				heightOffset,
			)
			rgba := image.NewRGBA(frame.Image.Bounds())
			draw.Draw(rgba, rgba.Bounds(), frame.Image, image.Point{}, draw.Src)

			err := imgtool.DrawTextOutline(rgba, &imgtool.TextOption{
				X:        10,
				Y:        55,
				FontData: font.ReadMicrosoftYaHei(),
				Size:     48,
				String:   tool.FormatDuration(frame.TimePointDuration),
			})
			if err != nil {
				tool.Logger.Error(err)
				tool.ErrorExit()
			}

			imgtool.DrawImageResize(
				backgroundImg,
				rgba,
				defaultFrameHeight,
				defaultFrameWidth,
				column*defaultFrameWidth,
				heightOffset,
			)
		}
	}

	return backgroundImg
}
