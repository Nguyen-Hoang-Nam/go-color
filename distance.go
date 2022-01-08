package gocolor

import "github.com/lucasb-eyer/go-colorful"

const (
	DistanceRgb = iota
	DistanceLab
	DistanceLuv
	DistanceCIE94
	DistanceCIEDE2000
)

func minDistanceRbg(color colorful.Color, limit int) int {
	minDistance := color.DistanceRgb(xterm256[0])
	minKey := 0
	for key, value := range xterm256 {
		if key == limit {
			break
		}

		valueDistance := color.DistanceRgb(value)
		if valueDistance < minDistance {
			minDistance = valueDistance
			minKey = key
		}
	}

	return minKey
}

func minDistanceLab(color colorful.Color, limit int) int {
	minDistance := color.DistanceLab(xterm256[0])
	minKey := 0
	for key, value := range xterm256 {
		if key == limit {
			break
		}

		valueDistance := color.DistanceLab(value)
		if valueDistance < minDistance {
			minDistance = valueDistance
			minKey = key
		}
	}

	return minKey
}

func minDistanceLuv(color colorful.Color, limit int) int {
	minDistance := color.DistanceLuv(xterm256[0])
	minKey := 0
	for key, value := range xterm256 {
		if key == limit {
			break
		}

		valueDistance := color.DistanceLuv(value)
		if valueDistance < minDistance {
			minDistance = valueDistance
			minKey = key
		}
	}

	return minKey
}

func minDistanceCIE94(color colorful.Color, limit int) int {
	minDistance := color.DistanceCIE94(xterm256[0])
	minKey := 0
	for key, value := range xterm256 {
		if key == limit {
			break
		}

		valueDistance := color.DistanceCIE94(value)
		if valueDistance < minDistance {
			minDistance = valueDistance
			minKey = key
		}
	}

	return minKey
}

func minDistanceCIEDE2000(color colorful.Color, limit int) int {
	minDistance := color.DistanceCIEDE2000(xterm256[0])
	minKey := 0
	for key, value := range xterm256 {
		if key == limit {
			break
		}

		valueDistance := color.DistanceCIEDE2000(value)
		if valueDistance < minDistance {
			minDistance = valueDistance
			minKey = key
		}
	}

	return minKey
}

func minDistance(color colorful.Color, distance int, limit int) int {
	switch distance {
	case DistanceRgb:
		return minDistanceRbg(color, limit)
	case DistanceLab:
		return minDistanceLab(color, limit)
	case DistanceLuv:
		return minDistanceLuv(color, limit)
	case DistanceCIE94:
		return minDistanceCIE94(color, limit)
	case DistanceCIEDE2000:
		return minDistanceCIEDE2000(color, limit)
	default:
		return 0
	}
}
