/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     app.go
 * +===============================================
 */
package http

import (
	"fmt"
	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/gorilla/mux"
	"net/http"
)

type HTTPHandlerFunc func(http.ResponseWriter, *http.Request, bh.Hive)
type HTTPHandler interface {
	Handler(http.ResponseWriter, *http.Request, bh.Hive)
}

type HTTPApp struct {
	App  bh.App
	Hive bh.Hive
}

func (h HTTPApp) HandleHTTPFunc(url string, handler HTTPHandlerFunc) *mux.Route {
	std_handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		handler(w, r, h.Hive)
	}
	return h.App.HandleHTTPFunc(url, std_handler)
}

func (h HTTPApp) HandleHTTP(url string, handler HTTPHandler) *mux.Route {
	std_handler := func(w http.ResponseWriter, r *http.Request) {
		handler.Handler(w, r, h.Hive)
	}
	return h.App.HandleHTTPFunc(url, std_handler)
}

func (h HTTPApp) DefaultHandle() *mux.Route {
	return h.HandleHTTPFunc("/{submodule}/{verb}", defaultHTTPHandler)
}

func NewHTTPApp(a bh.App, h bh.Hive) HTTPApp {
	return HTTPApp{
		App:  a,
		Hive: h,
	}
}
