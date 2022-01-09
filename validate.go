package gocolor

import "fmt"

func hsvValidate(hue float64, saturation float64, value float64) error {
	if hue < 0 || hue > 360 {
		return fmt.Errorf("Hue is out of range.\n")
	}

	if saturation < 0 || saturation > 1 {
		return fmt.Errorf("Saturation is out of range.\n")
	}

	if value < 0 || value > 1 {
		return fmt.Errorf("Value is out of range.\n")
	}

	return nil
}

func ansiValidate(ansiCode uint8) error {
	if ansiCode > 107 {
		return fmt.Errorf("Ansi code is out of range.\n")
	}

	return nil
}
