package intent

import (
	"github.com/elahejalalpour/beehive-netctrl/http"
	"github.com/elahejalalpour/beehive-netctrl/nom"
	bh "github.com/kandoo/beehive"
)


func RegisterIntent(h bh.Hive) {
	a := h.NewApp("intent")

	http.NewHTTPApp(a, h).DefaultHandle()
	a.Handle(http.HTTPRequest{}, )
}