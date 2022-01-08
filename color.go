package gocolor

import (
	"fmt"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

type ColorAttributes struct {
	attributes []string
}

const (
	escape      = "\x1b"
	fgColor256  = "[38;5"
	bgColor256  = "[48;5"
	fgTrueColor = "[38;2"
	bgTrueColor = "[48;2"
)

const (
	TrueColor = iota
	Color256
	Ansi
	NoColor
)

const (
	fg = iota
	bg
)

const (
	Reset     = escape + "[0m"
	Bold      = escape + "[1m"
	Dim       = escape + "[2m"
	Italic    = escape + "[3m"
	Underline = escape + "[4m"
	Reverse   = escape + "[7m"

	FgBlack   = escape + "[30m"
	FgRed     = escape + "[31m"
	FgGreen   = escape + "[32m"
	FgYellow  = escape + "[33m"
	FgBlue    = escape + "[34m"
	FgMagenta = escape + "[35m"
	FgCyan    = escape + "[36m"
	FgWhite   = escape + "[37m"

	BgBlack   = escape + "[40m"
	BgRed     = escape + "[41m"
	BgGreen   = escape + "[42m"
	BgYellow  = escape + "[43m"
	BgBlue    = escape + "[44m"
	BgMagenta = escape + "[45m"
	BgCyan    = escape + "[46m"
	BgWhite   = escape + "[47m"

	FgBrightBlack   = escape + "[90m"
	FgBrightRed     = escape + "[91m"
	FgBrightGreen   = escape + "[92m"
	FgBrightYellow  = escape + "[93m"
	FgBrightBlue    = escape + "[94m"
	FgBrightMagenta = escape + "[95m"
	FgBrightCyan    = escape + "[96m"
	FgBrightWhite   = escape + "[97m"

	BgBrightBlack   = escape + "[100m"
	BgBrightRed     = escape + "[101m"
	BgBrightGreen   = escape + "[102m"
	BgBrightYellow  = escape + "[103m"
	BgBrightBlue    = escape + "[104m"
	BgBrightMagenta = escape + "[105m"
	BgBrightCyan    = escape + "[106m"
	BgBrightWhite   = escape + "[107m"
)

func New(attributes ...string) *ColorAttributes {
	return &ColorAttributes{
		attributes: attributes,
	}
}

func (c *ColorAttributes) Add(attributes ...string) *ColorAttributes {
	c.attributes = append(c.attributes, attributes...)

	return c
}

func (c *ColorAttributes) Println(text string) {
	fmt.Printf("%s%s%s\n", strings.Join(c.attributes, ""), text, Reset)
}

func rgb(red uint8, green uint8, blue uint8, ground int) string {
	result := ""
	switch termColor() {
	case TrueColor:
		groundEscape := ""
		if ground == fg {
			groundEscape = fgTrueColor
		} else if ground == bg {
			groundEscape = bgTrueColor
		}

		result = fmt.Sprintf("%s%s;%d;%d;%dm", escape, groundEscape, red, green, blue)
	case Color256:
		currentColor := colorful.Color{
			R: float64(red) / 255, G: float64(green) / 255, B: float64(blue) / 255,
		}

		minDistance := currentColor.DistanceRgb(xterm256[0])
		minKey := 0
		for key, value := range xterm256 {
			valueDistance := currentColor.DistanceRgb(value)
			if valueDistance < minDistance {
				minDistance = valueDistance
				minKey = key
			}
		}

		groundEscape := ""
		if ground == fg {
			groundEscape = fgColor256
		} else if ground == bg {
			groundEscape = bgColor256
		}

		approximateColor := xterm256[minKey]
		approximateRed := int(approximateColor.R * 255)
		approximateGreen := int(approximateColor.G * 255)
		approximateBlue := int(approximateColor.B * 255)

		result = fmt.Sprintf("%s%s;%d;%d;%dm", escape, groundEscape, approximateRed, approximateGreen, approximateBlue)

	case Ansi:
		currentColor := colorful.Color{
			R: float64(red) / 255, G: float64(green) / 255, B: float64(blue) / 255,
		}

		minDistance := currentColor.DistanceRgb(xterm256[0])
		minKey := 0
		for key := 0; key < 16; key++ {
			valueDistance := currentColor.DistanceRgb(xterm256[key])
			if valueDistance < minDistance {
				minDistance = valueDistance
				minKey = key
			}
		}

		codeEscape := 30
		if ground == fg {
			if minKey < 8 {
				codeEscape = 30 + minKey
			} else {
				codeEscape = 82 + minKey
			}
		} else if ground == bg {
			if minKey < 8 {
				codeEscape = 40 + minKey
			} else {
				codeEscape = 92 + minKey
			}
		}

		result = fmt.Sprintf("%s[%dm", escape, codeEscape)
	case NoColor:
		result = ""
	default:
		result = ""
	}

	return result
}

func FgRGB(red uint8, green uint8, blue uint8) string {
	return rgb(red, green, blue, fg)
}

func BgRGB(red uint8, green uint8, blue uint8) string {
	return rgb(red, green, blue, bg)
}

func FgHSV(hue float64, saturation float64, value float64) string {
	color := colorful.Hsv(hue, saturation, value)

	return rgb(uint8(color.R*255), uint8(color.G*255), uint8(color.B*255), fg)
}

func BgHSV(hue float64, saturation float64, value float64) string {
	color := colorful.Hsv(hue, saturation, value)

	return rgb(uint8(color.R*255), uint8(color.G*255), uint8(color.B*255), bg)
}
