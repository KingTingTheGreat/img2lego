package handlers

import (
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
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
