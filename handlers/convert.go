package handlers

import (
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/kingtingthegreat/img2lego/converter"
	"github.com/kingtingthegreat/img2lego/util"
	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
)

const (
	maxUploadSize = 10 << 20 // 10 MB
	defaultWidth  = 100
)

func IsAppleHEIC(r io.Reader) (bool, error) {
	// HEIC brand is within first 12 bytes
	buf := make([]byte, 12)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return false, err
	}

	// bytes 4–7 must be "ftyp"
	if string(buf[4:8]) != "ftyp" {
		return false, nil
	}

	brand := string(buf[8:12])

	switch brand {
	case "heic", "heix", "hevc", "hevx", "mif1", "msf1":
		return true, nil
	default:
		return false, nil
	}
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

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
	img, _, err := image.Decode(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to read body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	img = resize.Resize(width, 0, img, resize.Lanczos3)

	// TODO clean up this heic check
	// check if image is heic
	heic, err := IsAppleHEIC(r.Body)
	if err != nil {
		log.Println(err)
		heic = false
	}

	// if heic, rotate 90 degrees clockwise
	if heic {
		b := img.Bounds()
		w, h := b.Dx(), b.Dy()

		dst := image.NewRGBA(image.Rect(0, 0, h, w))

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				dst.Set(h-1-y, x, img.At(b.Min.X+x, b.Min.Y+y))
			}
		}

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
