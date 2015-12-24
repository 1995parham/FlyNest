package intent

import (
	"github.com/elahejalalpour/beehive-netctrl/http"
	"github.com/elahejalalpour/beehive-netctrl/nom"
	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive-netctrl/discovery"
)


func RegisterIntent(h bh.Hive) {
	a := h.NewApp("intent")

	a.Handle(nom.LinkAdded{}, &discovery.GraphBuilderCentralized{})
	a.Handle(nom.LinkDeleted{}, &discovery.GraphBuilderCentralized{})

	http.NewHTTPApp(a, h).DefaultHandle()
//	a.Handle(http.HTTPRequest{}, &intentHandler{})
}