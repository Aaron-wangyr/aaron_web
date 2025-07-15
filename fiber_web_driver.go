package aaronweb

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type FiberWebDriver struct {
	app *fiber.App
}

// Start starts the web server
func (d *FiberWebDriver) Start() error {
	d.app = fiber.New()
	return nil
}

// Stop stops the web server
func (d *FiberWebDriver) Stop() error {
	if d.app != nil {
		return d.app.Shutdown()
	}
	return nil
}

func (d *FiberWebDriver) AddHandler(method string, path string, handler http.HandlerFunc) error {
	if d.app == nil {
		return errors.New("web driver not started")
	}

	switch method {
	case "GET":
		d.app.Get(path, adaptor.HTTPHandler(handler))
	case "POST":
		d.app.Post(path, adaptor.HTTPHandler(handler))
	case "PUT":
		d.app.Put(path, adaptor.HTTPHandler(handler))
	case "DELETE":
		d.app.Delete(path, adaptor.HTTPHandler(handler))
	default:
		return errors.New("unsupported HTTP method")
	}
	return nil
}

func (d *FiberWebDriver) RegisterMiddleware(middleware http.HandlerFunc) error {
	if d.app == nil {
		return errors.New("web driver not started")
	}

	d.app.Use(adaptor.HTTPHandler(middleware))
	return nil
}
