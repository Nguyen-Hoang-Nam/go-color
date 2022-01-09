package gocolor

import "github.com/lucasb-eyer/go-colorful"

func complementary(color colorful.Color) colorful.Color {
	return colorful.Color{
		R: 1 - color.R,
		G: 1 - color.G,
		B: 1 - color.B,
	}
}

func mapAnsiToColor256(ansiCode int) int {
	if ansiCode < 8 {
		return -1
	} else if ansiCode < 40 {
		return ansiCode - 30
	} else if ansiCode < 90 {
		return ansiCode - 40
	} else if ansiCode < 100 {
		return ansiCode - 90
	} else {
		return ansiCode - 100
	}
}

func getAnsiDisplay(ansiCode int) int {
	if ansiCode > 99 || (ansiCode < 90 && ansiCode > 39) {
		return bg
	} else if ansiCode < 8 {
		return -1
	}

	return fg
}
