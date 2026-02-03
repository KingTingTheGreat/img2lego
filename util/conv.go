package util

import (
	"strconv"
	"strings"

	"github.com/kyroy/kdtree/points"
)

type RGB struct {
	Data *points.Point3D
}

func HexToRGB(hexStr string) RGB {
	hexStr = strings.TrimPrefix(hexStr, "#")

	data := points.Point3D{}
	for i := 0; i < len(hexStr); i += 2 {
		value, _ := strconv.ParseFloat(hexStr[i:i+2], 16)
		switch i {
		case 0:
			data.X = value
		case 2:
			data.Y = value
		case 4:
			data.Z = value
		}
	}

	return RGB{
		Data: &data,
	}
}
