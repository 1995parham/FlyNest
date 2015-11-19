/*
 * In The Name Of God
 * ========================================
 * [] File Name : bh.go
 *
 * [] Creation Date : 19-11-2015
 *
 * [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Elahe Jalalpour.
 */
package discovery

import (
	"github.com/elahejalalpour/beehive-netctrl/nom"
	bh "github.com/kandoo/beehive"

	"time"
)

// RegisterDiscovery registers the handlers for topology discovery on the hive.
func RegisterDiscovery(h bh.Hive) {
	a := h.NewApp("discovery")
	a.Handle(nom.NodeJoined{}, &nodeJoinedHandler{})
	a.Handle(nom.NodeLeft{}, &nodeLeftHandler{})

	a.Handle(nom.PortUpdated{}, &portUpdateHandler{})
	// TODO(soheil): Handle PortRemoved.

	a.Handle(nom.PacketIn{}, &lldpPktInHandler{})

	a.Handle(nom.PacketIn{}, &arpPktInHandler{})

	a.Handle(NewLink{}, &newLinkHandler{})
	a.Handle(lldpTimeout{}, &timeoutHandler{})
	go func() {
		for {
			h.Emit(lldpTimeout{})
			time.Sleep(60 * time.Second)
		}
	}()
}
