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
	err        error
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
	Reset     = 0
	Bold      = 1
	Dim       = 2
	Italic    = 3
	Underline = 4
	Reverse   = 7

	FgBlack   = 30
	FgRed     = 31
	FgGreen   = 32
	FgYellow  = 33
	FgBlue    = 34
	FgMagenta = 35
	FgCyan    = 36
	FgWhite   = 37

	BgBlack   = 40
	BgRed     = 41
	BgGreen   = 42
	BgYellow  = 43
	BgBlue    = 44
	BgMagenta = 45
	BgCyan    = 46
	BgWhite   = 47

	FgBrightBlack   = 90
	FgBrightRed     = 91
	FgBrightGreen   = 92
	FgBrightYellow  = 93
	FgBrightBlue    = 94
	FgBrightMagenta = 95
	FgBrightCyan    = 96
	FgBrightWhite   = 97

	BgBrightBlack   = 100
	BgBrightRed     = 101
	BgBrightGreen   = 102
	BgBrightYellow  = 103
	BgBrightBlue    = 104
	BgBrightMagenta = 105
	BgBrightCyan    = 106
	BgBrightWhite   = 107
)

func New(attributes ...colorAttribute) *ColorAttributes {
	for _, attribute := range attributes {
		if attribute.err != nil {
			fmt.Printf("%v", attribute.err)
			return nil
		}
	}

	return &ColorAttributes{
		attributes: attributes,
		flag:       Auto,
		distance:   DistanceRgb,
	}
}

func AnsiEscape(escapeCode uint8) colorAttribute {
	err := ansiValidate(escapeCode)
	if err != nil {
		return colorAttribute{
			err: err,
		}
	}

	display := getAnsiDisplay(escapeCode)
	color256Code := mapAnsiToColor256(escapeCode)

	if color256Code == 255 {
		return colorAttribute{
			escapeCode: fmt.Sprintf("%s[%dm", escape, escapeCode),
			display:    display,
		}
	} else {
		return colorAttribute{
			escapeCode: fmt.Sprintf("%s[%dm", escape, escapeCode),
			display:    display,
			color:      xterm256[color256Code],
		}
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

func (c *ColorAttributes) AddComplementary() *ColorAttributes {
	numberAttribute := len(c.attributes)
	for i := numberAttribute - 1; i >= 0; i-- {
		attribute := c.attributes[i]
		if attribute.display == fg {
			complementaryAttribute := colorAttribute{
				display: bg,
				color:   complementary(attribute.color),
			}

			c.attributes = append(c.attributes, complementaryAttribute)

			break
		} else if attribute.display == bg {
			complementaryAttribute := colorAttribute{
				display: fg,
				color:   complementary(attribute.color),
			}

			c.attributes = append(c.attributes, complementaryAttribute)

			break
		}
	}

	return c
}

func (c *ColorAttributes) getAttributes() string {
	attributes := ""
	countFg := 0
	countBg := 0

	for i, attribute := range c.attributes {
		currentAttribute := attribute.escapeCode
		isAdd := false

		if attribute.display == fg {
			countFg++
			if countFg <= 1 {
				isAdd = true
			}
		} else if attribute.display == bg {
			countBg++
			if countFg <= 1 {
				isAdd = true
			}
		} else {
			isAdd = true
		}

		if isAdd {
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
	}

	return attributes
}

func (c *ColorAttributes) Println(text string) {
	fmt.Printf("%s%s%s[%dm\n", c.getAttributes(), text, escape, Reset)
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
	minKey := minDistance(color, distance, 256)

	groundEscape := ""
	if ground == fg {
		groundEscape = fgColor256
	} else if ground == bg {
		groundEscape = bgColor256
	}

	return fmt.Sprintf("%s%s;%dm", escape, groundEscape, minKey)
}

func ansi(color colorful.Color, ground int, distance int) string {
	minKey := minDistance(color, distance, 16)

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
		getTermColor()
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
	err := hsvValidate(hue, saturation, value)
	if err != nil {
		return colorAttribute{
			err: err,
		}
	}

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
	err := hsvValidate(hue, saturation, value)
	if err != nil {
		return colorAttribute{
			err: err,
		}
	}

	color := colorful.Hsv(hue, saturation, value)

	return colorAttribute{
		escapeCode: "",
		color: colorful.Color{
			R: color.R, G: color.G, B: color.B,
		},
		display: bg,
	}
}

func FgColor256(color256 uint8) colorAttribute {
	return colorAttribute{
		escapeCode: "",
		color:      xterm256[color256],
		display:    fg,
	}
}

func BgColor256(color256 uint8) colorAttribute {
	return colorAttribute{
		escapeCode: "",
		color:      xterm256[color256],
		display:    bg,
	}
}
