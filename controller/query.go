package controller

import (
	"github.com/1995parham/flynest/nom"
	bh "github.com/kandoo/beehive"
)

type queryHandler struct{}

func (h queryHandler) Rcv(msg bh.Msg, ctx bh.RcvContext) error {
	query := msg.Data().(nom.FlowStatsQuery)
	return sendToMaster(query, query.Node, ctx)
}

func (h queryHandler) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {
	return nodeDriversMap(msg.Data().(nom.FlowStatsQuery).Node)
}
