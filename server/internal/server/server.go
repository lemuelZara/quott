package server

import "net/http"

func NewWebApplication(router *http.ServeMux) *http.Server {
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &srv
}
