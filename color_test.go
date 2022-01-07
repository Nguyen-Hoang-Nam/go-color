package gocolor

import (
	"testing"
)

func TestColor(t *testing.T) {
	randomColor := New(FgRGB(103, 254, 201))
	randomColor.Println("Random color")

	ansiColor := New(FgRed)
	ansiColor.Println("Frontground red")

	combineAnsiColor := New(BgYellow, FgBlue)
	combineAnsiColor.Println("Combine Ansi color")

	reverseColor := New(BgYellow, FgBlue, Reverse)
	reverseColor.Println("Reverse color")
}
