package gocolor

import (
	"os"
	"runtime"
	"strings"

	"github.com/Nguyen-Hoang-Nam/go-color/terminal"
	"github.com/lucasb-eyer/go-colorful"
)

func isNoColorWindow() bool {
	if os.Getenv(term) == "dumb" {
		return true
	}

	if os.Getenv(noColor) != "" {
		return true
	}

	return false
}

func isNoColorLinux() bool {
	if os.Getenv(term) == "" || os.Getenv(term) == "dump" {
		return true
	}

	if os.Getenv("NO_COLOR") != "" {
		return true
	}

	return false
}

func isNoColor() bool {
	osType := runtime.GOOS

	switch osType {
	case "windows":
		return isNoColorWindow()

	case "linux":
		return isNoColorLinux()

	default:
		return true
	}
}

// Based on https://github.com/termstandard/colors
func isTrueColor() bool {
	colorTermEnv := os.Getenv(colorTerm)
	if colorTermEnv == "truecolor" || colorTermEnv == "24bit" {
		return true
	}

	termEnv := os.Getenv(term)
	switch termEnv {
	case
		"iterm",
		"linux-truecolor",
		"screen-truecolor",
		"tmux-truecolor",
		"xterm-truecolor",
		"xterm-kitty":
		return true
	}

	return strings.Contains(termEnv, "vte")
}

func isColor256() bool {
	return strings.Contains(os.Getenv(term), "256")
}

func getTerm() int {
	termEnv := os.Getenv(term)
	switch termEnv {
	case "xterm-kitty":
		return Kitty
	}

	return Default
}

func getTermColor() []colorful.Color {
	termName := getTerm()

	if termName == Kitty {
		colors, err := terminal.GetKittyColor()
		if err != nil {
			return terminal.Xterm256
		}

		return colors
	} else if termName == Alacritty {
		colors, err := terminal.GetAlacrittyColor()
		if err != nil {
			return terminal.Xterm256
		}

		return colors
	}

	return terminal.Xterm256
}

func termColor() int {
	if isNoColor() {
		return NoColor
	}

	// If terminal supports true color, then it also supports 256-colors
	if isTrueColor() {
		return TrueColor
	}

	// If terminal supports 256-colors, then is also supports 16-colors
	if isColor256() {
		return Color256
	}

	return Ansi
}
