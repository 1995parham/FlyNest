package intent

import (

	"github.com/elahejalalpour/beehive-netctrl/http"
	"github.com/elahejalalpour/beehive-netctrl/nom"
	bh "github.com/kandoo/beehive"
	"github.com/elahejalalpour/beehive-netctrl/discovery"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/golang/glog"

	"encoding/json"
	"fmt"
)

type intentHandler struct{}

type shortestPathData struct {
	from nom.UID
	to   nom.UID
}

func (h *intentHandler) Rcv(msg bh.Msg, ctx bh.RcvContext) error {
	hrq := msg.Data().(http.HTTPRequest)
	fmt.Println(hrq.AppName, hrq.Verb)
	if hrq.AppName == "intent" && hrq.Verb == "build" {
		fmt.Println(hrq.Data)
		spd := shortestPathData{}
		err := json.Unmarshal(hrq.Data, spd)
		if err != nil {
			glog.Errorf("Host list JSON marshaling: %v", err)
			return err
		}
		fmt.Println(discovery.ShortestPathCentralized(spd.from, spd.to, ctx))

		hrs := http.HTTPResponse{
			AppName: "host",
			Data: []byte{'A'},
		}

		err = ctx.Reply(msg, hrs)
		if err != nil {
			glog.Errorf("Replay error: %v", err)
			return err
		}
	}
	return nil

}

func (h *intentHandler) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {
	return bh.MappedCells{}
}