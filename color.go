package gocolor

import (
	"fmt"
	"strings"
)

type Color struct {
	attributes []string
}

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
	fmt.Printf("%s%s\n", strings.Join(c.attributes, ""), text)
}

func RGB(red int8, green int8, blue int8) string {
	return fmt.Sprintf("\x1B[38;2;%d;%d;%dm", red, green, blue)
}
