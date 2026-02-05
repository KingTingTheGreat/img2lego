package output

import (
	"fmt"

	"github.com/kingtingthegreat/img2lego/util"
)

func PrintRGBGrid(grid [][]util.RGB) {
	for _, row := range grid {
		for _, cell := range row {
			if cell.Data == nil {
				// fallback for missing data
				fmt.Print(" ")
				continue
			}

			r := clamp(int(cell.Data.X))
			g := clamp(int(cell.Data.Y))
			b := clamp(int(cell.Data.Z))

			// Truecolor foreground
			fmt.Printf("\x1b[38;2;%d;%d;%dm$\x1b[0m", r, g, b)
		}
		fmt.Println()
	}
}

func clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return v
}
