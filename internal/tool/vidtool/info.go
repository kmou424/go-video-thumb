package vidtool

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/mathutil"
	"github.com/kmou424/go-video-thumb/internal/tool"
	"github.com/kmou424/go-video-thumb/internal/tool/mathtool"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"path/filepath"
)

type Stream struct {
	Codec              string
	CodecLongName      string
	Width              int
	Height             int
	DisplayAspectRatio string
	Bitrate            string
	DurationSec        int
	Duration           string
	FrameRate          float64
	PixelFormat        string
}

type StreamInfo struct {
	Type               string `json:"codec_type"`
	Codec              string `json:"codec_name"`
	CodecLongName      string `json:"codec_long_name"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	ExpBitrate         string `json:"bit_rate"`
	ExpDuration        string `json:"duration"`
	ExpRFrameRate      string `json:"r_frame_rate"`
	PixelFormat        string `json:"pix_fmt"`
}

func (s *StreamInfo) AsStream() *Stream {
	durationSec := func() int {
		calculateInt, err := mathtool.CalculateInt(s.ExpDuration)
		if err != nil {
			tool.Logger.Error("an error occurred while calculating duration:", err)
			tool.ErrorExit()
		}
		return calculateInt
	}()
	return &Stream{
		Codec:              s.Codec,
		CodecLongName:      s.CodecLongName,
		Width:              s.Width,
		Height:             s.Height,
		DisplayAspectRatio: s.DisplayAspectRatio,
		Bitrate: func() string {
			cal, err := mathtool.CalculateFloat64(s.ExpBitrate)
			if err != nil {
				return "unknown"
			}

			return fmt.Sprintf("%s kbps", fmt.Sprintf("%g", cal/1000.0))
		}(),
		DurationSec: durationSec,
		Duration:    tool.FormatDuration(durationSec),
		FrameRate: func() float64 {
			if s.Type != "video" {
				return 0
			}
			calculateFloat64, err := mathtool.CalculateFloat64(s.ExpRFrameRate)
			if err != nil {
				return -1
			}
			return calculateFloat64
		}(),
		PixelFormat: s.PixelFormat,
	}
}

type FormatInfo struct {
	Size string `json:"size"`
}

type VideoInfo struct {
	StreamInfos []StreamInfo `json:"streams"`
	FormatInfo  `json:"format"`
	FileName    string `json:"-"`
}

func (v *VideoInfo) FormatSize() map[string]string {
	sizes := map[string]string{
		"KB": "",
		"MB": "",
		"GB": "",
		"TB": "",
	}
	sizesUnits := []string{"KB", "MB", "GB", "TB"}

	sizeBytes, err := mathutil.ToFloat(v.Size)
	if err != nil {
		tool.Logger.Error(err)
		tool.ErrorExit()
	}

	for _, unit := range sizesUnits {
		sizeBytes /= 1024
		sizes[unit] = fmt.Sprintf("%.3f", sizeBytes)
	}
	sizes["B"] = v.Size

	return sizes
}

func GetVideoInfo(file string) (*VideoInfo, error) {
	probe, err := ffmpeg.Probe(file)
	if err != nil {
		return nil, err
	}

	info := &VideoInfo{}
	err = json.Unmarshal([]byte(probe), info)
	if err != nil {
		return nil, err
	}

	info.FileName = filepath.Base(file)

	return info, nil
}
