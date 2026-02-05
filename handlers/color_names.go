package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kingtingthegreat/img2lego/converter"
)

type ColorName struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

func ColorNamesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	colorNames := make([]ColorName, len(converter.LEGO_COLORS))
	for i, c := range converter.LEGO_COLORS {
		colorNames[i] = ColorName{
			Color: c.Hex,
			Name:  c.Name,
		}
	}
	json.NewEncoder(w).Encode(colorNames)
}
