package router

import (
	"net/http"

	"github.com/kingtingthegreat/img2lego/handlers"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/health", handlers.HealthHandler)

	router.HandleFunc("GET /bg-colors", handlers.BGColorsHandler)
	router.HandleFunc("GET /color-names", handlers.ColorNamesHandler)

	router.HandleFunc("POST /convert", handlers.ConvertHandler)

	return router
}
