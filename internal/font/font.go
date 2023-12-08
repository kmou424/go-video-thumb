package font

import "embed"

//go:embed src
var fs embed.FS

func ReadMicrosoftYaHei() []byte {
	file, _ := fs.ReadFile("src/MicrosoftYaHei-Regular.ttf")
	return file
}
