package gocolor

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

const (
	term      = "TERM"
	noColor   = "NO_COLOR"
	colorTerm = "COLORTERM"
)

const (
	Kitty = iota
	Alacritty
	Default
)
