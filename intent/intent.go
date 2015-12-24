package intent

import (
	"github.com/elahejalalpour/beehive-netctrl/http"
	"github.com/elahejalalpour/beehive-netctrl/nom"
	bh "github.com/kandoo/beehive"
)

type intentHandler struct{}


func (h *intentHandler) Rcv(msg bh.Msg, ctx bh.RcvContext) error {

}

func (h *intentHandler) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {

}