package gocolor

import (
	"os"
	"runtime"
	"strings"
)

const (
	term      = "TERM"
	noColor   = "NO_COLOR"
	colorTerm = "COLORTERM"
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
		"xterm-truecolor":
		return true
	}

	if strings.Contains(termEnv, "vte") {
		return true
	}

	return false
}

func isColor256() bool {
	if strings.Contains(os.Getenv(term), "256") {
		return true
	}

	return false
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