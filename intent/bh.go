package intent

import (
	"github.com/1995parham/flynest/http"
	"github.com/1995parham/flynest/nom"
	bh "github.com/kandoo/beehive"
	"github.com/kandoo/flynest/discovery"
)


func RegisterIntent(h bh.Hive) {
	a := h.NewApp("intent")

	a.Handle(nom.LinkAdded{}, &discovery.GraphBuilderCentralized{})
	a.Handle(nom.LinkDeleted{}, &discovery.GraphBuilderCentralized{})

	http.NewHTTPApp(a, h).DefaultHandle()
	a.Handle(http.HTTPRequest{}, &intentHandler{})
}
