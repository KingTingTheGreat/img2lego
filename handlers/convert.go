package handlers

import (
	"bufio"
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/dolmen-go/kittyimg"
	"github.com/kingtingthegreat/img2lego/converter"
	"github.com/kingtingthegreat/img2lego/util"
	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
)

const (
	maxUploadSize = 10 << 20 // 10 MB
	defaultWidth  = 100
)

func IsAppleHEICHeader(peek []byte) bool {
	if len(peek) < 12 {
		return false
	}
	if string(peek[4:8]) != "ftyp" {
		return false
	}
	switch string(peek[8:12]) {
	case "heic", "heix", "hevc", "hevx", "mif1", "msf1":
		return true
	default:
		return false
	}
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	br := bufio.NewReader(r.Body)
	peek, _ := br.Peek(12)

	var width uint
	qWidth, err := strconv.Atoi(r.URL.Query().Get("width"))
	if err != nil || qWidth <= 0 {
		width = defaultWidth
	} else {
		width = uint(qWidth)
	}

	// check that body is not empty
	if r.ContentLength == 0 {
		http.Error(w, "Missing image", http.StatusBadRequest)
		return
	}

	// limit the total size of the request body
	// r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	// raw image in body
	img, _, err := image.Decode(br)
	if err != nil {
		log.Println("error decoding image", err)
		http.Error(w, "failed to read body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	img = resize.Resize(width, 0, img, resize.Lanczos3)

	// TODO clean up this heic check
	// check if image is heic
	heic := IsAppleHEICHeader(peek)
	// if heic, rotate 90 degrees clockwise
	if heic {
		log.Println("Rotating HEIC image")
		b := img.Bounds()
		w, h := b.Dx(), b.Dy()

		dst := image.NewRGBA(image.Rect(0, 0, h, w))

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				dst.Set(h-1-y, x, img.At(b.Min.X+x, b.Min.Y+y))
			}
		}

		kittyimg.Fprintln(os.Stdout, dst) // to avoid memory leak in kittyimg
		kittyimg.Fprintln(os.Stdout, img) // to avoid memory leak in kittyimg

		img = dst
	}

	res := converter.ConvertImage(img)

	convertedImage := make([][]string, len(res))
	for i, row := range res {
		convertedRow := make([]string, len(row))
		for j, cell := range row {
			convertedRow[j] = util.RGBToHex(cell)
		}
		convertedImage[i] = convertedRow
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(convertedImage)
}
