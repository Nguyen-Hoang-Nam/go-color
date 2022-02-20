package gocolor

import (
	"testing"
)

func TestColor(t *testing.T) {
	rgbColor := FgRGB(70, 165, 224)

	randomColor := New(rgbColor)
	randomColor.Println("Random color")

	ansiColor := New(AnsiEscape(FgRed))
	ansiColor.Println("New(AnsiEscape(FgRed))")

	complementaryAnsiColor := New(AnsiEscape(FgRed)).AddComplementary()
	complementaryAnsiColor.Println("New(AnsiEscape(FgRed)).AddComplementary()")

	combineAnsiColor := New(AnsiEscape(BgYellow), AnsiEscape(FgBlue))
	combineAnsiColor.Println("New(AnsiEscape(BgYellow), AnsiEscape(FgBlue))")

	reverseColor := New(AnsiEscape(BgYellow), AnsiEscape(FgBlue), AnsiEscape(Reverse))
	reverseColor.Println("New(AnsiEscape(BgYellow), AnsiEscape(FgBlue), AnsiEscape(Reverse))")

	hsvColor := New(FgHSV(120, 1, 0.5))
	hsvColor.Println("New(FgHSV(120, 1, 0.5))")

	flagColor256 := New(rgbColor).SetFlag(Color256)
	flagColor256.Println("New(rgbColor).SetFlag(Color256)")

	flagAnsi := New(rgbColor).SetFlag(Ansi)
	flagAnsi.Println("New(rgbColor).SetFlag(Ansi)")

	flagAnsiDistanceLab := New(rgbColor).SetFlag(Ansi).SetDistance(DistanceLab)
	flagAnsiDistanceLab.Println("New(rgbColor).SetFlag(Ansi).SetDistance(DistanceLab)")

	distanceLab := New(rgbColor).SetFlag(Color256).SetDistance(DistanceLab)
	distanceLab.Println("New(rgbColor).SetFlag(Color256).SetDistance(DistanceLab)")

	distanceCIE94 := New(rgbColor).SetFlag(Color256).SetDistance(DistanceCIE94)
	distanceCIE94.Println("New(rgbColor).SetFlag(Color256).SetDistance(DistanceCIE94)")

	importConfiguration := New().SetFallbackFromPath("./configuration.json").Add(FgColor256(24))
	importConfiguration.Println("New(FgColor256(24))")
}
