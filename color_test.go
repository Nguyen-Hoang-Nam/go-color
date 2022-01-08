package gocolor

import (
	"testing"
)

func TestColor(t *testing.T) {
	rgbColor := FgRGB(70, 165, 224)

	randomColor := New(rgbColor)
	randomColor.Println("Random color")

	ansiColor := New(AnsiEscape(FgRed))
	ansiColor.Println("Frontground red")

	complementaryAnsiColor := New(rgbColor).AddComplementary()
	complementaryAnsiColor.Println("Complementary background")

	combineAnsiColor := New(AnsiEscape(BgYellow), AnsiEscape(FgBlue))
	combineAnsiColor.Println("Combine Ansi color")

	reverseColor := New(AnsiEscape(BgYellow), AnsiEscape(FgBlue), AnsiEscape(Reverse))
	reverseColor.Println("Reverse color")

	hsvColor := New(FgHSV(120, 1, 0.5))
	hsvColor.Println("HSV color")

	flagColor256 := New(rgbColor).SetFlag(Color256)
	flagColor256.Println("Flag 256")

	flagAnsi := New(rgbColor).SetFlag(Ansi)
	flagAnsi.Println("Flag ansi")

	flagAnsiDistanceLab := New(rgbColor).SetFlag(Ansi).SetDistance(DistanceLab)
	flagAnsiDistanceLab.Println("Flag ansi with Lab distance")

	distanceLab := New(rgbColor).SetFlag(Color256).SetDistance(DistanceLab)
	distanceLab.Println("Distance lab")

	distanceCIE94 := New(rgbColor).SetFlag(Color256).SetDistance(DistanceCIE94)
	distanceCIE94.Println("Distance CIE94")
}
