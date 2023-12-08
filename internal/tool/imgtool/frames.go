package imgtool

import (
	"bytes"
	"github.com/kmou424/go-video-thumb/internal/tool"
	"github.com/kmou424/go-video-thumb/internal/tool/vidtool"
	"image"
	"image/png"
)

type Frame struct {
	Image             image.Image
	TimePointDuration int
	Count             int
}

func calculateTimeMarks(duration int, imageColumns int, imageRows int) []int {
	totalFrames := imageColumns * imageRows
	timeMarks := make([]int, totalFrames)

	frameInterval := duration / (totalFrames + 1)

	remainder := duration % (totalFrames + 1)

	for i := 0; i < totalFrames; i++ {
		timeMarks[i] = frameInterval * (i + 1)

		if remainder > 0 {
			timeMarks[i] += remainder
			remainder--
		}
	}

	return timeMarks
}

func ExtractFrames(inputFile string, vStream vidtool.Stream, imageColumns int, imageRows int) (frames []*Frame) {
	for i, timeMark := range calculateTimeMarks(vStream.DurationSec, imageColumns, imageRows) {
		tool.Logger.Info("extracting frame #", i+1)
		frameBuffer, err := vidtool.CutFrameFromVideo(inputFile, timeMark)
		if err != nil {
			tool.Logger.Error(err)
			return
		}
		img, err := png.Decode(bytes.NewReader(frameBuffer.Bytes()))
		if err != nil {
			tool.Logger.Error(err)
			return
		}
		frames = append(frames, &Frame{
			Image:             img,
			TimePointDuration: timeMark,
			Count:             i + 1,
		})
	}
	return
}
