package controller

import (
	"github.com/elahejalalpour/beehive-netctrl/nom"
	bh "github.com/kandoo/beehive"
)

type pktOutHandler struct{}

func (h pktOutHandler) Rcv(msg bh.Msg, ctx bh.RcvContext) error {
	pkt := msg.Data().(nom.PacketOut)
	return sendToMaster(pkt, pkt.Node, ctx)
}

func (h pktOutHandler) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {
	return nodeDriversMap(msg.Data().(nom.PacketOut).Node)
}
