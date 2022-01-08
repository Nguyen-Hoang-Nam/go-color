package gocolor

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

type ColorAttributes struct {
	attributes []colorAttribute
	flag       int
	distance   int
}

type colorAttribute struct {
	escapeCode string
	color      colorful.Color
	display    int
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
	Auto
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

func New(attributes ...colorAttribute) *ColorAttributes {
	return &ColorAttributes{
		attributes: attributes,
		flag:       Auto,
		distance:   DistanceRgb,
	}
}

func AnsiEscape(escapeCode string) colorAttribute {
	return colorAttribute{
		escapeCode: escapeCode,
	}
}

func (c *ColorAttributes) SetFlag(flag int) *ColorAttributes {
	c.flag = flag

	return c
}

func (c *ColorAttributes) SetDistance(distance int) *ColorAttributes {
	c.distance = distance

	return c
}

func (c *ColorAttributes) Add(attributes ...colorAttribute) *ColorAttributes {
	c.attributes = append(c.attributes, attributes...)

	return c
}

func (c *ColorAttributes) Println(text string) {
	attributes := ""
	for i, attribute := range c.attributes {
		currentAttribute := attribute.escapeCode
		if currentAttribute == "" {
			escapeCode := ""
			isColor := true

			if c.flag == Auto {
				escapeCode = rgb(attribute.color, attribute.display, c.distance)
			} else if c.flag == TrueColor {
				escapeCode = trueColor(attribute.color, attribute.display)
			} else if c.flag == Color256 {
				escapeCode = color256(attribute.color, attribute.display, c.distance)
			} else if c.flag == Ansi {
				escapeCode = ansi(attribute.color, attribute.display, c.distance)
			} else {
				isColor = false
			}

			if isColor {
				attributes += escapeCode

				attribute.escapeCode = escapeCode
				c.attributes[i] = attribute
			}
		} else {
			attributes += currentAttribute
		}
	}

	fmt.Printf("%s%s%s\n", attributes, text, Reset)
}

func trueColor(color colorful.Color, ground int) string {
	groundEscape := ""
	if ground == fg {
		groundEscape = fgTrueColor
	} else if ground == bg {
		groundEscape = bgTrueColor
	}

	return fmt.Sprintf("%s%s;%d;%d;%dm", escape, groundEscape, uint8(color.R*255), uint8(color.G*255), uint8(color.B*255))
}

func color256(color colorful.Color, ground int, distance int) string {
	minKey := minDistance(color, distance)

	groundEscape := ""
	if ground == fg {
		groundEscape = fgColor256
	} else if ground == bg {
		groundEscape = bgColor256
	}

	return fmt.Sprintf("%s%s;%dm", escape, groundEscape, minKey)
}

func ansi(color colorful.Color, ground int, distance int) string {
	minKey := minDistance(color, distance)

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

	return fmt.Sprintf("%s[%dm", escape, codeEscape)
}

func rgb(color colorful.Color, ground int, distance int) string {
	switch termColor() {
	case TrueColor:
		return trueColor(color, ground)

	case Color256:
		return color256(color, ground, distance)

	case Ansi:
		return ansi(color, ground, distance)

	case NoColor:
		return ""

	default:
		return ""
	}
}

func FgRGB(red uint8, green uint8, blue uint8) colorAttribute {
	return colorAttribute{
		escapeCode: "",
		color: colorful.Color{
			R: float64(red) / 255, G: float64(green) / 255, B: float64(blue) / 255,
		},
		display: fg,
	}
}

func BgRGB(red uint8, green uint8, blue uint8) colorAttribute {
	return colorAttribute{
		escapeCode: "",
		color: colorful.Color{
			R: float64(red) / 255, G: float64(green) / 255, B: float64(blue) / 255,
		},
		display: bg,
	}
}

func FgHSV(hue float64, saturation float64, value float64) colorAttribute {
	color := colorful.Hsv(hue, saturation, value)

	return colorAttribute{
		escapeCode: "",
		color: colorful.Color{
			R: color.R, G: color.G, B: color.B,
		},
		display: fg,
	}
}

func BgHSV(hue float64, saturation float64, value float64) colorAttribute {
	color := colorful.Hsv(hue, saturation, value)

	return colorAttribute{
		escapeCode: "",
		color: colorful.Color{
			R: color.R, G: color.G, B: color.B,
		},
		display: bg,
	}
}
