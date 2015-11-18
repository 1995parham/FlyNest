/*
 * In The Name Of God
 * ========================================
 * [] File Name : bh.go
 *
 * [] Creation Date : 18-11-2015
 *
 * [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Elahe Jalalpour.
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

	app, ok := vars["app"]
	if !ok {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	fmt.Println(app)

	feature, ok := vars["feature"]
	if !ok {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	fmt.Println(feature)

	if r.Method == "POST" {
		fmt.Println("Hello from POST message")
	}

	if r.Method == "GET" {
		fmt.Println("Hello from GET message")
	}

	creq := ConfigRequest{
		AppName:         app,
		RequiredFeature: feature,
	}

	cres, err := bh.Sync(context.TODO(), creq)
	if err == nil && cres != nil {
		w.Header().Set("Server", "Beehive-netctrl-Config-Server")
		w.Write(cres.(ConfigResponse).ResponseData)
		return
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Errorf("ConfigHTTPHandler: %v\n", err)
		return
	}
}

func StartConfig(h bh.Hive) error {
	app := h.NewApp("ConfigApp", bh.Persistent(1))
	app.HandleHTTP("/{app}/{feature}", &ConfigHTTPHandler{})

	return nil
}
