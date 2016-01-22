package intent

import (

	"github.com/1995parham/flynest/http"
	"github.com/1995parham/flynest/nom"
	bh "github.com/kandoo/beehive"
	"github.com/1995parham/flynest/discovery"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/golang/glog"

	"encoding/json"
	"fmt"
)

type intentHandler struct{}

type shortestPathData struct {
	From nom.UID    `json:"from"`
	To   nom.UID    `json:"to"`
}

func (h *intentHandler) Rcv(msg bh.Msg, ctx bh.RcvContext) error {
	hrq := msg.Data().(http.HTTPRequest)
	if hrq.AppName == "intent" && hrq.Verb == "build" {
		spd := shortestPathData{}
		err := json.Unmarshal(hrq.Data, &spd)
		if err != nil {
			glog.Errorf("Host list JSON unmarshaling: %v", err)
			return err
		}
		fmt.Println(spd)
		fmt.Println(discovery.ShortestPathCentralized(spd.From, spd.To, ctx))

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
	return bh.RuntimeMap(h.Rcv)(msg, ctx)
}
