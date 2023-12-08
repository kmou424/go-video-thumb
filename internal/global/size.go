package global

func DefaultFrameWidth(height int, width int) int {
	if height > width {
		// DefaultFrameWidthPortrait
		return 360
	}
	// DefaultFrameWidth
	return 480
}
