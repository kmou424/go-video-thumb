package vidtool

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"golang.org/x/net/context"
	"io"
	"os"
)

func CutFrameFromVideo(videoPath string, time int) (*bytes.Buffer, error) {
	ffmpeg.LogCompiledCommand = false
	buf := bytes.NewBuffer(nil)

	stream := ffmpeg.Input(videoPath, ffmpeg.KwArgs{"ss": time}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "png"}).
		WithOutput(buf, os.Stdout)
	stream.Context = context.WithValue(stream.Context, "Stderr", io.Discard)

	err := stream.Run()

	if err != nil {
		return nil, err
	}

	return buf, nil
}
