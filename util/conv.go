package util

import (
	"fmt"
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
		value, _ := strconv.ParseUint(hexStr[i:i+2], 16, 8)

		switch i {
		case 0:
			data.X = float64(value)
		case 2:
			data.Y = float64(value)
		case 4:
			data.Z = float64(value)
		}
	}

	return RGB{
		Data: &data,
	}
}

func RGBToHex(rgb RGB) string {
	hex := fmt.Sprintf(
		"#%02X%02X%02X",
		clampToU8(rgb.Data.X),
		clampToU8(rgb.Data.Y),
		clampToU8(rgb.Data.Z),
	)

	return strings.ToUpper(hex)
}

// clamp converts a float-ish channel into 0..255 safely.
func clampToU8(v float64) uint8 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v + 0.5) // round to nearest
}
