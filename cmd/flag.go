package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"github.com/kmou424/go-video-thumb/internal/global"
	"github.com/kmou424/go-video-thumb/internal/tool"
	"os"
)

var (
	inputFile    string
	imageColumns int
	imageRows    int
	fontSize     int
	noHeader     bool

	version bool
)

func parseArgs() {
	flag.StringVar(&inputFile, "i", "", "input video file")
	flag.IntVar(&imageColumns, "cols", 5, "columns of thumbnails grid")
	flag.IntVar(&imageRows, "rows", 4, "rows of thumbnails grid")
	flag.IntVar(&fontSize, "font-size", 20, "font size of header text")
	flag.BoolVar(&noHeader, "no-header", false, "do not generate header")
	flag.BoolVar(&version, "version", false, "show version")

	flag.Parse()

	if version {
		fmt.Printf(`video-thumb v%s
Copyright (c) 2023 kmou424
`, global.Version)
		os.Exit(0)
	}

	err := checkArgs()
	if err != nil {
		tool.Logger.Fatal("an error occurred:", err)
		tool.FatalExit()
	}
}

func checkArgs() error {
	// handle windows drag event
	if len(os.Args) == 2 {
		if !strutil.IsStartOf(os.Args[1], "-") && fsutil.FileExists(os.Args[1]) {
			inputFile = os.Args[1]
		}
	}

	if inputFile == "" || !fsutil.FileExists(inputFile) {
		return errors.New("invalid input file")
	}

	return nil
}
