/*
 * In The Name Of God
 * ========================================
 * [] File Name : bh.go
 *
 * [] Creation Date : 18-11-2015
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Parham Alvani.
 */
package config

import (
	"fmt"
	"net/http"

	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive/Godeps/_workspace/src/golang.org/x/net/context"
)

type ConfigHTTPHandler struct{}

func (h *ConfigHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	fmt.Println(url.RequestURI())

	if r.Method == "POST" {
		fmt.Println("Hello from POST message")
	}

	if r.Method == "GET" {
		fmt.Println("Hello from GET message")
	}

	creq := ConfigRequest{
		AppName: url.RequestURI(),
	}

	cres, err := bh.Sync(context.TODO(), creq)
	if err == nil {
		w.Header().Set("Server", "Beehive-netctrl-Config-Server")
		w.Write(cres.(ConfigResponse).ResponseData)
	}
}

func StartConfig(h bh.Hive) {
	app := h.NewApp("hello-world", bh.Persistent(1))
	app.HandleHTTP("/{name}", &ConfigHTTPHandler{})
}
