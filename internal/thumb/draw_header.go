package thumb

import (
	"fmt"
	"github.com/kmou424/go-video-thumb/internal/font"
	"github.com/kmou424/go-video-thumb/internal/global"
	"github.com/kmou424/go-video-thumb/internal/tool"
	"github.com/kmou424/go-video-thumb/internal/tool/imgtool"
	"image"
)

func drawHeaderImage() *image.RGBA {
	sizeMap := videoInfo.FormatSize()
	headerStrings := []string{
		fmt.Sprintf("Filename: %s", videoInfo.FileName),
		fmt.Sprintf("Size: %s MB (%s bytes)", sizeMap["MB"], sizeMap["B"]),
		fmt.Sprintf("Resolution: %dx%d (%s)", vStream.Width, vStream.Height, vStream.DisplayAspectRatio),
		fmt.Sprintf("FrameRate: %g fps", vStream.FrameRate),
		fmt.Sprintf("Video Codec: %s [%s], pixfmt: %s", vStream.Codec, vStream.CodecLongName, vStream.PixelFormat),
		fmt.Sprintf("Audio Codec: %s [%s]", aStream.Codec, aStream.CodecLongName),
		fmt.Sprintf("Bitrate: %s", vStream.Bitrate),
		fmt.Sprintf("Duration: %s", vStream.Duration),
	}

	perLineHeight := global.FontSize + 3
	offsetY := 10
	backgroundImg := imgtool.DrawWhiteImage(len(headerStrings)*perLineHeight+offsetY*2, defaultFrameWidth*global.ImageColumns)

	startX := 10
	startY := offsetY - 5
	for _, headerString := range headerStrings {
		startY += perLineHeight
		err := imgtool.DrawText(backgroundImg, &imgtool.TextOption{
			X:        startX,
			Y:        startY,
			FontData: font.ReadMicrosoftYaHei(),
			Size:     float64(global.FontSize),
			String:   headerString,
		})
		if err != nil {
			tool.Logger.Error(err)
			tool.ErrorExit()
		}
	}

	return backgroundImg
}
