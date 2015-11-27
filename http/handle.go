/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     handle.go
 * +===============================================
 */
package http

import (
	"fmt"
	"io/ioutil"
	"net/http"

	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/kandoo/beehive/Godeps/_workspace/src/golang.org/x/net/context"
)

func defaultHTTPHandler(w http.ResponseWriter, r *http.Request, h bh.Hive) {
	w.Header().Set("Server", "Beehive-netctrl-HTTP-Server")

	vars := mux.Vars(r)

	submodule, ok := vars["submodule"]
	if !ok {
		/* This should not happend :) */
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	verb, ok := vars["verb"]
	if !ok {
		/* This should not happend :) */
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	creq := HTTPRequest{
		AppName: submodule,
		Verb:    verb,
	}

	// Read content data if avaiable :)
	if r.ContentLength > 0 {
		data, err := ioutil.ReadAll(r.Body)
		if err == nil {
			creq.Data = data
		}
	}

	cres, err := h.Sync(context.TODO(), creq)
	if err == nil && cres != nil {
		w.Write(cres.(HTTPResponse).Data)
		w.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Errorf("defaultTTPHandler: %v\n", err)
		return
	}

}
