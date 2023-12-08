package main

import (
	"fmt"
	"github.com/kmou424/go-video-thumb/internal/font"
	"github.com/kmou424/go-video-thumb/internal/global"
	"github.com/kmou424/go-video-thumb/internal/tool"
	"github.com/kmou424/go-video-thumb/internal/tool/imgtool"
	"github.com/kmou424/go-video-thumb/internal/tool/vidtool"
	"image"
	"image/draw"
	"path/filepath"
)

var (
	videoInfo vidtool.VideoInfo
	vStream   vidtool.Stream
	aStream   vidtool.Stream
)

var (
	defaultFrameWidth  int
	defaultFrameHeight int
)

func drawVideoThumb() {
	info, err := vidtool.GetVideoInfo(inputFile)
	if err != nil {
		tool.Logger.Error(err)
		tool.ErrorExit()
	}
	videoInfo = *info
	for _, streamInfo := range videoInfo.StreamInfos {
		switch streamInfo.Type {
		case "video":
			vStream = *streamInfo.AsStream()
		case "audio":
			aStream = *streamInfo.AsStream()
		}
	}

	frames := imgtool.ExtractFrames(inputFile, vStream, imageColumns, imageRows)

	defaultFrameWidth = global.DefaultFrameWidth(vStream.Height, vStream.Width)
	defaultFrameHeight = int(float64(defaultFrameWidth) / float64(vStream.Width) * float64(vStream.Height))

	headerImage := drawHeaderImage()
	thumbnailImage := drawThumbnailImage(frames)

	var images []image.Image
	if !noHeader {
		images = append(images, headerImage)
	}
	images = append(images, thumbnailImage)

	combinedImage, err := imgtool.ConcatImages(images...)
	if err != nil {
		tool.Logger.Error(err)
		tool.ErrorExit()
	}

	imageFileName := inputFile
	imageFileName = imageFileName[:len(imageFileName)-len(filepath.Ext(imageFileName))]
	err = imgtool.SaveImage(combinedImage, fmt.Sprintf("%s.png", imageFileName))
	if err != nil {
		tool.Logger.Error(err)
		tool.ErrorExit()
	}
}

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

	perLineHeight := fontSize + 3
	offsetY := 10
	backgroundImg := imgtool.DrawWhiteImage(len(headerStrings)*perLineHeight+offsetY*2, defaultFrameWidth*imageColumns)

	startX := 10
	startY := offsetY - 5
	for _, headerString := range headerStrings {
		startY += perLineHeight
		err := imgtool.DrawText(backgroundImg, &imgtool.TextOption{
			X:        startX,
			Y:        startY,
			FontData: font.ReadMicrosoftYaHei(),
			Size:     float64(fontSize),
			String:   headerString,
		})
		if err != nil {
			tool.Logger.Error(err)
			tool.ErrorExit()
		}
	}

	return backgroundImg
}

func drawThumbnailImage(frames []*imgtool.Frame) *image.RGBA {
	backgroundImg := imgtool.DrawWhiteImage(defaultFrameHeight*imageRows, defaultFrameWidth*imageColumns)

	for row := 0; row < imageRows; row++ {
		heightOffset := defaultFrameHeight * row
		for column := 0; column < imageColumns; column++ {
			frame := frames[row*imageColumns+column]
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
