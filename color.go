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
	TrueColor = iota
	Color256
	Ansi
	NoColor
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
	fmt.Printf("%s%s\x1b[0m\n", strings.Join(c.attributes, ""), text)
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

func RGB(red uint8, green uint8, blue uint8) string {
	return fmt.Sprintf("\x1B[38;2;%d;%d;%dm", red, green, blue)
}
