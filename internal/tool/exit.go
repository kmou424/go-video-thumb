package tool

import "os"

func FatalExit() {
	os.Exit(-1)
}

func ErrorExit() {
	os.Exit(1)
}
