package global

import (
	"flag"
	"fmt"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"os"
)

var (
	Input        string
	ImageColumns int
	ImageRows    int
	FontSize     int
	NoHeader     bool

	Version bool
)

var InputType string

func ParseArgs() {
	flag.StringVar(&Input, "i", "", "input video file/directory")
	flag.IntVar(&ImageColumns, "cols", 5, "columns of thumbnails grid")
	flag.IntVar(&ImageRows, "rows", 4, "rows of thumbnails grid")
	flag.IntVar(&FontSize, "font-size", 20, "font size of header text")
	flag.BoolVar(&NoHeader, "no-header", false, "do not generate header")
	flag.BoolVar(&Version, "version", false, "show version")

	flag.Parse()
}

func CheckArgs() error {
	// handle windows drag event
	if len(os.Args) == 2 {
		if !strutil.IsStartOf(os.Args[1], "-") {
			Input = os.Args[1]
		}
	}

	switch {
	case fsutil.IsFile(Input):
		InputType = "file"
	case fsutil.IsDir(Input):
		InputType = "dir"
	default:
		return fmt.Errorf("invalid input: %s", Input)
	}

	return nil
}
