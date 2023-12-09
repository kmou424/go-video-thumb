package main

import (
	"fmt"
	"github.com/kmou424/go-video-thumb/internal/global"
	"github.com/kmou424/go-video-thumb/internal/thumb"
	"github.com/kmou424/go-video-thumb/internal/tool"
	"os"
)

var BuildVersion string = "unknown"

func main() {
	global.ParseArgs()

	if global.Version {
		fmt.Printf(`video-thumb v%s
Copyright (c) 2023 kmou424
`, BuildVersion)
		os.Exit(0)
	}

	err := global.CheckArgs()
	if err != nil {
		tool.Logger.Fatal("an error occurred:", err)
		tool.FatalExit()
	}

	thumb.Draw()
}
