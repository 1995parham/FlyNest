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
	"io/ioutil"
	"net/http"

	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/kandoo/beehive/Godeps/_workspace/src/golang.org/x/net/context"
)

func StartConfig(h bh.Hive) error {
	app := h.NewApp("ConfigApp", bh.Persistent(1))

	configHTTPHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "Beehive-netctrl-Config-Server")

		vars := mux.Vars(r)

		app, ok := vars["app"]
		if !ok {
			/* This should not happend :) */
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		feature, ok := vars["feature"]
		if !ok {
			/* This should not happend :) */
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		creq := ConfigRequest{
			AppName:         app,
			RequiredFeature: feature,
		}

		if r.ContentLength > 0 {
			data, err := ioutil.ReadAll(r.Body)
			if err == nil {
				creq.RequiredData = data
			}
		}

		cres, err := h.Sync(context.TODO(), creq)
		if err == nil && cres != nil {
			w.Write(cres.(ConfigResponse).ResponseData)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Errorf("ConfigHTTPHandler: %v\n", err)
			return
		}
	}

	app.HandleHTTPFunc("/{app}/{feature}", configHTTPHandler)

	return nil
}
