package gocolor

import "github.com/lucasb-eyer/go-colorful"

const (
	Xterm0 = 0
	Xterm1 = 128 / 255
	Xterm2 = 95 / 255
	Xterm3 = 135 / 255
	Xterm4 = 175 / 255
	Xterm5 = 215 / 255
	Xterm6 = 1
)

// Base on XTerm color
var Xterm256 = []colorful.Color{
	{R: Xterm0, G: Xterm0, B: Xterm0},
	{R: Xterm1, G: Xterm0, B: Xterm0},
	{R: Xterm0, G: Xterm1, B: Xterm0},
	{R: Xterm1, G: Xterm1, B: Xterm0},
	{R: Xterm0, G: Xterm0, B: Xterm1},
	{R: Xterm1, G: Xterm0, B: Xterm1},
	{R: 192 / 255, G: 192 / 255, B: 192 / 255},
	{R: Xterm1, G: Xterm1, B: Xterm1},
	{R: Xterm6, G: Xterm0, B: Xterm0},
	{R: Xterm0, G: Xterm6, B: Xterm0},
	{R: Xterm6, G: Xterm6, B: Xterm0},
	{R: Xterm0, G: Xterm0, B: Xterm6},
	{R: Xterm6, G: Xterm0, B: Xterm6},
	{R: Xterm0, G: Xterm6, B: Xterm6},
	{R: Xterm6, G: Xterm6, B: Xterm6},

	{R: Xterm0, G: Xterm0, B: Xterm0},
	{R: Xterm0, G: Xterm0, B: Xterm2},
	{R: Xterm0, G: Xterm0, B: Xterm3},
	{R: Xterm0, G: Xterm0, B: Xterm4},
	{R: Xterm0, G: Xterm0, B: Xterm5},
	{R: Xterm0, G: Xterm0, B: Xterm6},

	{R: Xterm0, G: Xterm2, B: Xterm0},
	{R: Xterm0, G: Xterm2, B: Xterm2},
	{R: Xterm0, G: Xterm2, B: Xterm3},
	{R: Xterm0, G: Xterm2, B: Xterm4},
	{R: Xterm0, G: Xterm2, B: Xterm5},
	{R: Xterm0, G: Xterm2, B: Xterm6},

	{R: Xterm0, G: Xterm3, B: Xterm0},
	{R: Xterm0, G: Xterm3, B: Xterm2},
	{R: Xterm0, G: Xterm3, B: Xterm3},
	{R: Xterm0, G: Xterm3, B: Xterm4},
	{R: Xterm0, G: Xterm3, B: Xterm5},
	{R: Xterm0, G: Xterm3, B: Xterm6},

	{R: Xterm0, G: Xterm4, B: Xterm0},
	{R: Xterm0, G: Xterm4, B: Xterm2},
	{R: Xterm0, G: Xterm4, B: Xterm3},
	{R: Xterm0, G: Xterm4, B: Xterm4},
	{R: Xterm0, G: Xterm4, B: Xterm5},
	{R: Xterm0, G: Xterm4, B: Xterm6},

	{R: Xterm0, G: Xterm5, B: Xterm0},
	{R: Xterm0, G: Xterm5, B: Xterm2},
	{R: Xterm0, G: Xterm5, B: Xterm3},
	{R: Xterm0, G: Xterm5, B: Xterm4},
	{R: Xterm0, G: Xterm5, B: Xterm5},
	{R: Xterm0, G: Xterm5, B: Xterm6},

	{R: Xterm0, G: Xterm6, B: Xterm0},
	{R: Xterm0, G: Xterm6, B: Xterm2},
	{R: Xterm0, G: Xterm6, B: Xterm3},
	{R: Xterm0, G: Xterm6, B: Xterm4},
	{R: Xterm0, G: Xterm6, B: Xterm5},
	{R: Xterm0, G: Xterm6, B: Xterm6},

	{R: Xterm2, G: Xterm0, B: Xterm0},
	{R: Xterm2, G: Xterm0, B: Xterm2},
	{R: Xterm2, G: Xterm0, B: Xterm3},
	{R: Xterm2, G: Xterm0, B: Xterm4},
	{R: Xterm2, G: Xterm0, B: Xterm5},
	{R: Xterm2, G: Xterm0, B: Xterm6},

	{R: Xterm2, G: Xterm2, B: Xterm0},
	{R: Xterm2, G: Xterm2, B: Xterm2},
	{R: Xterm2, G: Xterm2, B: Xterm3},
	{R: Xterm2, G: Xterm2, B: Xterm4},
	{R: Xterm2, G: Xterm2, B: Xterm5},
	{R: Xterm2, G: Xterm2, B: Xterm6},

	{R: Xterm2, G: Xterm3, B: Xterm0},
	{R: Xterm2, G: Xterm3, B: Xterm2},
	{R: Xterm2, G: Xterm3, B: Xterm3},
	{R: Xterm2, G: Xterm3, B: Xterm4},
	{R: Xterm2, G: Xterm3, B: Xterm5},
	{R: Xterm2, G: Xterm3, B: Xterm6},

	{R: Xterm2, G: Xterm4, B: Xterm0},
	{R: Xterm2, G: Xterm4, B: Xterm2},
	{R: Xterm2, G: Xterm4, B: Xterm3},
	{R: Xterm2, G: Xterm4, B: Xterm4},
	{R: Xterm2, G: Xterm4, B: Xterm5},
	{R: Xterm2, G: Xterm4, B: Xterm6},

	{R: Xterm2, G: Xterm5, B: Xterm0},
	{R: Xterm2, G: Xterm5, B: Xterm2},
	{R: Xterm2, G: Xterm5, B: Xterm3},
	{R: Xterm2, G: Xterm5, B: Xterm4},
	{R: Xterm2, G: Xterm5, B: Xterm5},
	{R: Xterm2, G: Xterm5, B: Xterm6},

	{R: Xterm2, G: Xterm6, B: Xterm0},
	{R: Xterm2, G: Xterm6, B: Xterm2},
	{R: Xterm2, G: Xterm6, B: Xterm3},
	{R: Xterm2, G: Xterm6, B: Xterm4},
	{R: Xterm2, G: Xterm6, B: Xterm5},
	{R: Xterm2, G: Xterm6, B: Xterm6},

	{R: Xterm3, G: Xterm0, B: Xterm0},
	{R: Xterm3, G: Xterm0, B: Xterm2},
	{R: Xterm3, G: Xterm0, B: Xterm3},
	{R: Xterm3, G: Xterm0, B: Xterm4},
	{R: Xterm3, G: Xterm0, B: Xterm5},
	{R: Xterm3, G: Xterm0, B: Xterm6},

	{R: Xterm3, G: Xterm2, B: Xterm0},
	{R: Xterm3, G: Xterm2, B: Xterm2},
	{R: Xterm3, G: Xterm2, B: Xterm3},
	{R: Xterm3, G: Xterm2, B: Xterm4},
	{R: Xterm3, G: Xterm2, B: Xterm5},
	{R: Xterm3, G: Xterm2, B: Xterm6},

	{R: Xterm3, G: Xterm3, B: Xterm0},
	{R: Xterm3, G: Xterm3, B: Xterm2},
	{R: Xterm3, G: Xterm3, B: Xterm3},
	{R: Xterm3, G: Xterm3, B: Xterm4},
	{R: Xterm3, G: Xterm3, B: Xterm5},
	{R: Xterm3, G: Xterm3, B: Xterm6},

	{R: Xterm3, G: Xterm4, B: Xterm0},
	{R: Xterm3, G: Xterm4, B: Xterm2},
	{R: Xterm3, G: Xterm4, B: Xterm3},
	{R: Xterm3, G: Xterm4, B: Xterm4},
	{R: Xterm3, G: Xterm4, B: Xterm5},
	{R: Xterm3, G: Xterm4, B: Xterm6},

	{R: Xterm3, G: Xterm5, B: Xterm0},
	{R: Xterm3, G: Xterm5, B: Xterm2},
	{R: Xterm3, G: Xterm5, B: Xterm3},
	{R: Xterm3, G: Xterm5, B: Xterm4},
	{R: Xterm3, G: Xterm5, B: Xterm5},
	{R: Xterm3, G: Xterm5, B: Xterm6},

	{R: Xterm3, G: Xterm6, B: Xterm0},
	{R: Xterm3, G: Xterm6, B: Xterm2},
	{R: Xterm3, G: Xterm6, B: Xterm3},
	{R: Xterm3, G: Xterm6, B: Xterm4},
	{R: Xterm3, G: Xterm6, B: Xterm5},
	{R: Xterm3, G: Xterm6, B: Xterm6},

	{R: Xterm4, G: Xterm0, B: Xterm0},
	{R: Xterm4, G: Xterm0, B: Xterm2},
	{R: Xterm4, G: Xterm0, B: Xterm3},
	{R: Xterm4, G: Xterm0, B: Xterm4},
	{R: Xterm4, G: Xterm0, B: Xterm5},
	{R: Xterm4, G: Xterm0, B: Xterm6},

	{R: Xterm4, G: Xterm2, B: Xterm0},
	{R: Xterm4, G: Xterm2, B: Xterm2},
	{R: Xterm4, G: Xterm2, B: Xterm3},
	{R: Xterm4, G: Xterm2, B: Xterm4},
	{R: Xterm4, G: Xterm2, B: Xterm5},
	{R: Xterm4, G: Xterm2, B: Xterm6},

	{R: Xterm4, G: Xterm3, B: Xterm0},
	{R: Xterm4, G: Xterm3, B: Xterm2},
	{R: Xterm4, G: Xterm3, B: Xterm3},
	{R: Xterm4, G: Xterm3, B: Xterm4},
	{R: Xterm4, G: Xterm3, B: Xterm5},
	{R: Xterm4, G: Xterm3, B: Xterm6},

	{R: Xterm4, G: Xterm4, B: Xterm0},
	{R: Xterm4, G: Xterm4, B: Xterm2},
	{R: Xterm4, G: Xterm4, B: Xterm3},
	{R: Xterm4, G: Xterm4, B: Xterm4},
	{R: Xterm4, G: Xterm4, B: Xterm5},
	{R: Xterm4, G: Xterm4, B: Xterm6},

	{R: Xterm4, G: Xterm5, B: Xterm0},
	{R: Xterm4, G: Xterm5, B: Xterm2},
	{R: Xterm4, G: Xterm5, B: Xterm3},
	{R: Xterm4, G: Xterm5, B: Xterm4},
	{R: Xterm4, G: Xterm5, B: Xterm5},
	{R: Xterm4, G: Xterm5, B: Xterm6},

	{R: Xterm4, G: Xterm6, B: Xterm0},
	{R: Xterm4, G: Xterm6, B: Xterm2},
	{R: Xterm4, G: Xterm6, B: Xterm3},
	{R: Xterm4, G: Xterm6, B: Xterm4},
	{R: Xterm4, G: Xterm6, B: Xterm5},
	{R: Xterm4, G: Xterm6, B: Xterm6},

	{R: Xterm5, G: Xterm0, B: Xterm0},
	{R: Xterm5, G: Xterm0, B: Xterm2},
	{R: Xterm5, G: Xterm0, B: Xterm3},
	{R: Xterm5, G: Xterm0, B: Xterm4},
	{R: Xterm5, G: Xterm0, B: Xterm5},
	{R: Xterm5, G: Xterm0, B: Xterm6},

	{R: Xterm5, G: Xterm2, B: Xterm0},
	{R: Xterm5, G: Xterm2, B: Xterm2},
	{R: Xterm5, G: Xterm2, B: Xterm3},
	{R: Xterm5, G: Xterm2, B: Xterm4},
	{R: Xterm5, G: Xterm2, B: Xterm5},
	{R: Xterm5, G: Xterm2, B: Xterm6},

	{R: Xterm5, G: Xterm3, B: Xterm0},
	{R: Xterm5, G: Xterm3, B: Xterm2},
	{R: Xterm5, G: Xterm3, B: Xterm3},
	{R: Xterm5, G: Xterm3, B: Xterm4},
	{R: Xterm5, G: Xterm3, B: Xterm5},
	{R: Xterm5, G: Xterm3, B: Xterm6},

	{R: Xterm5, G: Xterm4, B: Xterm0},
	{R: Xterm5, G: Xterm4, B: Xterm2},
	{R: Xterm5, G: Xterm4, B: Xterm3},
	{R: Xterm5, G: Xterm4, B: Xterm4},
	{R: Xterm5, G: Xterm4, B: Xterm5},
	{R: Xterm5, G: Xterm4, B: Xterm6},

	{R: Xterm5, G: Xterm5, B: Xterm0},
	{R: Xterm5, G: Xterm5, B: Xterm2},
	{R: Xterm5, G: Xterm5, B: Xterm3},
	{R: Xterm5, G: Xterm5, B: Xterm4},
	{R: Xterm5, G: Xterm5, B: Xterm5},
	{R: Xterm5, G: Xterm5, B: Xterm6},

	{R: Xterm5, G: Xterm6, B: Xterm0},
	{R: Xterm5, G: Xterm6, B: Xterm2},
	{R: Xterm5, G: Xterm6, B: Xterm3},
	{R: Xterm5, G: Xterm6, B: Xterm4},
	{R: Xterm5, G: Xterm6, B: Xterm5},
	{R: Xterm5, G: Xterm6, B: Xterm6},

	{R: Xterm6, G: Xterm0, B: Xterm0},
	{R: Xterm6, G: Xterm0, B: Xterm2},
	{R: Xterm6, G: Xterm0, B: Xterm3},
	{R: Xterm6, G: Xterm0, B: Xterm4},
	{R: Xterm6, G: Xterm0, B: Xterm5},
	{R: Xterm6, G: Xterm0, B: Xterm6},

	{R: Xterm6, G: Xterm2, B: Xterm0},
	{R: Xterm6, G: Xterm2, B: Xterm2},
	{R: Xterm6, G: Xterm2, B: Xterm3},
	{R: Xterm6, G: Xterm2, B: Xterm4},
	{R: Xterm6, G: Xterm2, B: Xterm5},
	{R: Xterm6, G: Xterm2, B: Xterm6},

	{R: Xterm6, G: Xterm3, B: Xterm0},
	{R: Xterm6, G: Xterm3, B: Xterm2},
	{R: Xterm6, G: Xterm3, B: Xterm3},
	{R: Xterm6, G: Xterm3, B: Xterm4},
	{R: Xterm6, G: Xterm3, B: Xterm5},
	{R: Xterm6, G: Xterm3, B: Xterm6},

	{R: Xterm6, G: Xterm4, B: Xterm0},
	{R: Xterm6, G: Xterm4, B: Xterm2},
	{R: Xterm6, G: Xterm4, B: Xterm3},
	{R: Xterm6, G: Xterm4, B: Xterm4},
	{R: Xterm6, G: Xterm4, B: Xterm5},
	{R: Xterm6, G: Xterm4, B: Xterm6},

	{R: Xterm6, G: Xterm5, B: Xterm0},
	{R: Xterm6, G: Xterm5, B: Xterm2},
	{R: Xterm6, G: Xterm5, B: Xterm3},
	{R: Xterm6, G: Xterm5, B: Xterm4},
	{R: Xterm6, G: Xterm5, B: Xterm5},
	{R: Xterm6, G: Xterm5, B: Xterm6},

	{R: Xterm6, G: Xterm6, B: Xterm0},
	{R: Xterm6, G: Xterm6, B: Xterm2},
	{R: Xterm6, G: Xterm6, B: Xterm3},
	{R: Xterm6, G: Xterm6, B: Xterm4},
	{R: Xterm6, G: Xterm6, B: Xterm5},
	{R: Xterm6, G: Xterm6, B: Xterm6},

	{R: 8 / 255, G: 8 / 255, B: 8 / 255},
	{R: 18 / 255, G: 18 / 255, B: 18 / 255},
	{R: 28 / 255, G: 28 / 255, B: 28 / 255},
	{R: 38 / 255, G: 38 / 255, B: 38 / 255},
	{R: 48 / 255, G: 48 / 255, B: 48 / 255},
	{R: 58 / 255, G: 58 / 255, B: 58 / 255},
	{R: 68 / 255, G: 68 / 255, B: 68 / 255},
	{R: 78 / 255, G: 78 / 255, B: 78 / 255},
	{R: 88 / 255, G: 88 / 255, B: 88 / 255},
	{R: 98 / 255, G: 98 / 255, B: 98 / 255},
	{R: 108 / 255, G: 108 / 255, B: 108 / 255},
	{R: 118 / 255, G: 118 / 255, B: 118 / 255},
	{R: 128 / 255, G: 128 / 255, B: 128 / 255},
	{R: 138 / 255, G: 138 / 255, B: 138 / 255},
	{R: 148 / 255, G: 148 / 255, B: 148 / 255},
	{R: 158 / 255, G: 158 / 255, B: 158 / 255},
	{R: 168 / 255, G: 168 / 255, B: 168 / 255},
	{R: 178 / 255, G: 178 / 255, B: 178 / 255},
	{R: 188 / 255, G: 188 / 255, B: 188 / 255},
	{R: 198 / 255, G: 198 / 255, B: 198 / 255},
	{R: 208 / 255, G: 208 / 255, B: 208 / 255},
	{R: 218 / 255, G: 218 / 255, B: 218 / 255},
	{R: 228 / 255, G: 228 / 255, B: 228 / 255},
	{R: 238 / 255, G: 238 / 255, B: 238 / 255},
}
