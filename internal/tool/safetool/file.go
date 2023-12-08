package safetool

import (
	"github.com/kmou424/go-video-thumb/internal/tool"
	"os"
)

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		tool.Logger.Warn("close file failed: %s", err.Error())
	}
}
