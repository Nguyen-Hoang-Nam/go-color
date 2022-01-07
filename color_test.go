package gocolor

import (
	"testing"
)

func TestColor(t *testing.T) {
	c := New(RGB(103, 254, 201))
	c.Println("Hello world!!!")
}
