package converter

import (
	"image"

	"github.com/kingtingthegreat/img2lego/util"
	"github.com/kyroy/kdtree/points"
)

func ConvertImage(img image.Image) [][]util.RGB {
	l := InitLookup(LEGO_COLORS)

	b := img.Bounds()

	output := make([][]util.RGB, b.Dy())

	i := 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		row := make([]util.RGB, b.Dx())
		j := 0
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			closest := l.FindMostSimilar(util.RGB{
				Data: &points.Point3D{X: float64(r >> 8), Y: float64(g >> 8), Z: float64(b >> 8)},
			})

			row[j] = closest
			j++
		}

		output[i] = row
		i++
	}

	return output
}
