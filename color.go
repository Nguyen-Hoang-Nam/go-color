package gocolor

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Color struct {
	attributes []string
}

const (
	escape     = "\x1b"
	fgAnsi     = "[38;5"
	bgAnsi     = "[48;5"
	fgColor256 = "[38;2"
	bgColor256 = "[48;2"
)

const (
	TrueColor = iota
	Color256
	Ansi
	NoColor
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

func New(attributes ...string) *Color {
	c := &Color{
		attributes: make([]string, 0),
	}

	c.Add(attributes...)
	return c
}

func (c *Color) Add(attributes ...string) *Color {
	c.attributes = append(c.attributes, attributes...)

	return c
}

func (c *Color) Println(text string) {
	fmt.Printf("%s%s%s[0m\n", strings.Join(c.attributes, ""), text, escape)
}

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
	case "Linux":
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

func FgRGB(red uint8, green uint8, blue uint8) string {
	return fmt.Sprintf("%s%s;%d;%d;%dm", escape, fgColor256, red, green, blue)
}

func BgRGB(red uint8, green uint8, blue uint8) string {
	return fmt.Sprintf("%s%s;%d;%d;%dm", escape, bgColor256, red, green, blue)
}
