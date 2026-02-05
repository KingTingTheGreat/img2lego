package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kingtingthegreat/img2lego/converter"
)

func BGColorsHandler(
	w http.ResponseWriter, r *http.Request,
) {
	bgColors := make([]string, len(converter.BG_COLORS))
	for i, c := range converter.BG_COLORS {
		bgColors[i] = c.Hex
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bgColors)
}
