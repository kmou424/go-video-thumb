package tool

import (
	"fmt"
	"time"
)

func FormatDuration(durationSec int) string {
	duration := time.Duration(durationSec) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
