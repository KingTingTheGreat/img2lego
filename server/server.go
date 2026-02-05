package server

import (
	"net/http"

	"github.com/kingtingthegreat/img2lego/middleware"
	"github.com/kingtingthegreat/img2lego/router"
)

func Server() *http.Server {
	router := router.Router()

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.CORS(middleware.Logger(router)),
	}

	return &server
}
