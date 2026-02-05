package converter

import (
	"image"

	"github.com/kingtingthegreat/img2lego/util"
	"github.com/kyroy/kdtree/points"
)

func ConvertImage(img image.Image) [][]util.RGB {
	l := InitLookup(LEGO_COLORS)

	bounds := img.Bounds()

	output := make([][]util.RGB, bounds.Dy())

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		row := make([]util.RGB, bounds.Dx())
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			closest := l.FindMostSimilar(util.RGB{
				Data: &points.Point3D{
					X: float64(uint8(r >> 8)),
					Y: float64(uint8(g >> 8)),
					Z: float64(uint8(b >> 8))},
			})

			row[x-bounds.Min.X] = closest
		}

		output[y-bounds.Min.Y] = row
	}

	return output
}
