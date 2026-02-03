package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/dolmen-go/kittyimg"
	"github.com/kingtingthegreat/img2lego/converter"
	"github.com/kingtingthegreat/img2lego/util"
	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
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

func main() {
	// open image
	f, err := os.Open("LightsaberCSWE.webp")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	// scale image with converter
	img = resize.Resize(100, 0, img, resize.Lanczos3)
	kittyimg.Fprintln(os.Stdout, img)

	// convert image
	output := converter.ConvertImage(img)

	// print out converted image
	PrintRGBGrid(output)
}
