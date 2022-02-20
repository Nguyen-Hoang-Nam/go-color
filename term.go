package gocolor

import (
	"os"
	"runtime"
	"strings"

	"github.com/Nguyen-Hoang-Nam/go-color/terminal"
	"github.com/lucasb-eyer/go-colorful"
)

var cacheColors []colorful.Color = nil

func isNoColorWindow() bool {
	if os.Getenv(TERM) == "dumb" {
		return true
	}

	if os.Getenv(NO_COLOR) != "" {
		return true
	}

	return false
}

func isNoColorLinux() bool {
	if os.Getenv(TERM) == "" || os.Getenv(TERM) == "dump" {
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
	colorTermEnv := os.Getenv(COLOR_TERM)
	if colorTermEnv == "truecolor" || colorTermEnv == "24bit" {
		return true
	}

	termEnv := os.Getenv(TERM)
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
	return strings.Contains(os.Getenv(TERM), "256")
}

func getTerm() int {
	termEnv := os.Getenv(TERM)
	switch termEnv {
	case "xterm-kitty":
		return Kitty
	}

	return Default
}

func getTermColor() []colorful.Color {
	if cacheColors != nil {
		return cacheColors
	}

	termName := getTerm()

	if termName == Kitty {
		colors, err := terminal.GetKittyColor()
		if err != nil {
			return terminal.Xterm256
		}

		cacheColors = colors
		return colors
	} else if termName == Alacritty {
		colors, err := terminal.GetAlacrittyColor()
		if err != nil {
			return terminal.Xterm256
		}

		cacheColors = colors
		return colors
	}

	cacheColors = terminal.Xterm256
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
