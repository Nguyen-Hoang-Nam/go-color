package gocolor

import (
	"fmt"
	"os"
	"runtime"
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

// Credit https://github.com/fatih/color
func New(attributes ...string) *ColorAttributes {
	c := &ColorAttributes{
		attributes: make([]string, 0),
	}

	c.Add(attributes...)
	return c
}

func (c *ColorAttributes) Add(attributes ...string) *ColorAttributes {
	c.attributes = append(c.attributes, attributes...)

	return c
}

func (c *ColorAttributes) Println(text string) {
	fmt.Printf("%s%s%s%s\n", strings.Join(c.attributes, ""), text, escape, Reset)
}

// Inspirate by https://github.com/BurntSushi/termcolor
func allowColor() int {
	osType := runtime.GOOS

	term := os.Getenv("TERM")
	switch osType {
	case "windows":
		if term == "dumb" {
			return NoColor
		}

		if os.Getenv("NO_COLOR") != "" {
			return NoColor
		}

	case "linux":
		if term == "" || term == "dump" {
			return NoColor
		}

		if os.Getenv("NO_COLOR") != "" {
			return NoColor
		}

	default:
		return NoColor
	}

	colorTerm := os.Getenv("COLORTERM")
	if colorTerm == "truecolor" || colorTerm == "24bit" {
		return TrueColor
	}

	if strings.Contains(term, "256") {
		return Color256
	} else if term == "iterm" || term == "iterm2" {
		return Color256
	} else {
		return Ansi
	}
}

func rgb(red uint8, green uint8, blue uint8, ground int) string {
	result := ""
	switch allowColor() {
	case TrueColor:
		groundEscape := ""
		if ground == fg {
			groundEscape = fgTrueColor
		} else if ground == bg {
			groundEscape = bgTrueColor
		}

		result = fmt.Sprintf("%s%s;%d;%d;%dm", escape, groundEscape, red, green, blue)
	case Color256:
		result = ""
	case Ansi:
		result = ""
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
