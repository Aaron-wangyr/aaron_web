package aaronweb

import (
	"net/http"
)

type WebEngine interface {
	// Start the web server
	Start() error
	// Stop the web server
	Stop() error

	AddHandler(method string, path string, handler http.HandlerFunc) error

	RegisterMiddleware(middleware http.HandlerFunc) error
}
