package gocolor

import (
	"testing"
)

func TestColor(t *testing.T) {
	randomColor := New(FgRGB(103, 254, 201))
	randomColor.Println("Random color")

	ansiColor := New(AnsiEscape(FgRed))
	ansiColor.Println("Frontground red")

	combineAnsiColor := New(AnsiEscape(BgYellow), AnsiEscape(FgBlue))
	combineAnsiColor.Println("Combine Ansi color")

	reverseColor := New(AnsiEscape(BgYellow), AnsiEscape(FgBlue), AnsiEscape(Reverse))
	reverseColor.Println("Reverse color")

	hsvColor := New(FgHSV(120, 1, 0.5))
	hsvColor.Println("HSV color")

	flagColor256 := New(FgRGB(103, 254, 201)).SetFlag(Color256)
	flagColor256.Println("Flag 256")
}
