package terminal

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

func extractColor(content string) map[int]colorful.Color {
	outSlice := strings.Split(content, "\n")

	pattern := regexp.MustCompile(`color\d`)

	colorMap := map[int]colorful.Color{}

	for _, line := range outSlice {
		matched := pattern.MatchString(line)

		if matched {
			lineSlice := strings.Fields(line)
			numStr := lineSlice[0][5:]

			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			c, err := colorful.Hex(lineSlice[1])
			if err != nil {
				continue
			}

			colorMap[num] = c
		}
	}

	return colorMap
}

func getColorRemotControlProtocol() (map[int]colorful.Color, error) {
	cmd := exec.Command("kitty", "@", "get-colors")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	outStr := out.String()
	return extractColor(outStr), nil
}

func checkFileExit(filePath string) bool {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return true
	}

	return false
}

// TODO Recuring parse Kitty configuration to list of colors
func readConfiguration(filePath string) (map[int]colorful.Color, error) {
	if checkFileExit(filePath) {
		configuration, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		configurationStr := string(configuration)
		return extractColor(configurationStr), nil
	}

	return nil, errors.New("File not exist")
}

func getColorFromConfiguration() (map[int]colorful.Color, error) {
	KITTY_CONFIG_DIRECTORY := os.Getenv("KITTY_CONFIG_DIRECTORY")
	if KITTY_CONFIG_DIRECTORY != "" {
		if colors, err := readConfiguration(KITTY_CONFIG_DIRECTORY); err == nil {
			return colors, nil
		}
	}

	XDG_CONFIG_HOME := os.Getenv("XDG_CONFIG_HOME")
	if XDG_CONFIG_HOME != "" {
		if colors, err := readConfiguration(XDG_CONFIG_HOME + "/kitty/kitty.conf"); err == nil {
			return colors, nil
		}
	}

	HOME := os.Getenv("HOME")
	if colors, err := readConfiguration(HOME + "/.config/kitty/kitty.conf"); err == nil {
		return colors, nil
	}

	XDG_CONFIG_DIRS := os.Getenv("XDG_CONFIG_DIRS")
	if XDG_CONFIG_DIRS != "" {
		if colors, err := readConfiguration(XDG_CONFIG_DIRS + "/kitty/kitty.conf"); err == nil {
			return colors, nil
		}
	}

	return nil, errors.New("Can not get color")
}

// Run really slow
func GetKittyColor() (map[int]colorful.Color, error) {
	colors, err := getColorFromConfiguration()
	if err != nil {
		colors, err = getColorRemotControlProtocol()

		if err != nil {
			return nil, err
		}

		return colors, nil
	}

	return colors, nil
}
