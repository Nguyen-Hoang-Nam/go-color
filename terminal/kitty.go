package terminal

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

// Run really slow
func getColorRemotControlProtocol() ([]colorful.Color, error) {
	cmd := exec.Command("kitty", "@", "get-colors")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	outStr := out.String()
	outSlice := strings.Split(outStr, "\n")

	pattern := regexp.MustCompile(`color\d`)

	var colors []colorful.Color
	for _, line := range outSlice {
		matched := pattern.MatchString(line)

		if matched {
			lineSlice := strings.Fields(line)
			// numStr := lineSlice[0][5:]

			// num, err := strconv.Atoi(numStr)
			// if err != nil {
			// 	continue
			// }

			c, err := colorful.Hex(lineSlice[1])
			if err != nil {
				continue
			}

			// TODO use hasmap to get only avaiable colors
			// colors[num] = c
			colors = append(colors, c)
		}
	}

	return colors, nil
}

// TODO Recuring parse Kitty configuration to list of colors
func getColorFromConfiguration() ([]colorful.Color, error) {
	return nil, errors.New("Can not get color")
}

func GetKittyColor() ([]colorful.Color, error) {
	colors, err := getColorRemotControlProtocol()
	if err != nil {
		colors, err = getColorFromConfiguration()

		if err != nil {
			return nil, err
		}

		return colors, nil
	}

	return colors, nil
}
