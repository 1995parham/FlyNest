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
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/kandoo/beehive/Godeps/_workspace/src/golang.org/x/net/context"
)

type ConfigHTTPHandler struct{}

func (h *ConfigHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]

	if !ok {
		fmt.Println("Invalid request")
	}
	fmt.Println(name)

	if r.Method == "POST" {
		fmt.Println("Hello from POST message")
	}

	if r.Method == "GET" {
		fmt.Println("Hello from GET message")
	}

	creq := ConfigRequest{
		AppName: name,
	}

	cres, err := bh.Sync(context.TODO(), creq)
	if err == nil {
		w.Header().Set("Server", "Beehive-netctrl-Config-Server")
		w.Write(cres.(ConfigResponse).ResponseData)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartConfig(h bh.Hive) {
	app := h.NewApp("ConfigApp", bh.Persistent(1))
	app.HandleHTTP("/{name}", &ConfigHTTPHandler{})
}
