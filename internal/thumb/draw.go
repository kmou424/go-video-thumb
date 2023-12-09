package thumb

import (
	"fmt"
	"github.com/gookit/goutil/fsutil"
	"github.com/kmou424/go-video-thumb/internal/global"
	"github.com/kmou424/go-video-thumb/internal/tool"
	"github.com/kmou424/go-video-thumb/internal/tool/imgtool"
	"github.com/kmou424/go-video-thumb/internal/tool/vidtool"
	"image"
	"io/fs"
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

func Draw() {
	switch global.InputType {
	case "file":
		info, err := vidtool.GetVideoInfo(global.Input)
		if err != nil {
			tool.Logger.Error(err)
			tool.ErrorExit()
		}
		videoInfo = *info
		drawVideoThumb(global.Input)
	case "dir":
		err := fsutil.WalkDir(global.Input, func(path string, _ fs.DirEntry, err error) error {
			if err != nil {
				tool.Logger.Error("an error occurred while walking the directory:", err)
			}
			log := fmt.Sprintf("processing: %s", path)
			info, err := vidtool.GetVideoInfo(path)
			if err != nil {
				tool.Logger.Info(log, "not a video")
				return nil
			}
			videoInfo = *info
			tool.Logger.Info(log)
			drawVideoThumb(path)
			return nil
		})
		if err != nil {
			tool.Logger.Error(err)
			tool.ErrorExit()
		}
	}
}

func drawVideoThumb(videoFile string) {
	for _, streamInfo := range videoInfo.StreamInfos {
		switch streamInfo.Type {
		case "video":
			vStream = *streamInfo.AsStream()
		case "audio":
			aStream = *streamInfo.AsStream()
		}
	}

	frames := imgtool.ExtractFrames(videoFile, vStream, global.ImageColumns, global.ImageRows)

	defaultFrameWidth = global.DefaultFrameWidth(vStream.Height, vStream.Width)
	defaultFrameHeight = int(float64(defaultFrameWidth) / float64(vStream.Width) * float64(vStream.Height))

	headerImage := drawHeaderImage()
	thumbnailImage := drawThumbnailImage(frames)

	var images []image.Image
	if !global.NoHeader {
		images = append(images, headerImage)
	}
	images = append(images, thumbnailImage)

	combinedImage, err := imgtool.ConcatImages(images...)
	if err != nil {
		tool.Logger.Error(err)
		tool.ErrorExit()
	}

	imageFileName := videoInfo.FileName
	imageFileName = imageFileName[:len(imageFileName)-len(filepath.Ext(imageFileName))]
	err = imgtool.SaveImage(combinedImage, fmt.Sprintf("%s.png", imageFileName))
	if err != nil {
		tool.Logger.Error(err)
		tool.ErrorExit()
	}
}
