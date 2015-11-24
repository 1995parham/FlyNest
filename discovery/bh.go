/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 24-11-2015
 * |
 * | File Name:     bh.go
 * +===============================================
 */
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

// RegisterDiscovery registers the discovery module for topology discovery on the hive.
// you can use it's REST API in order to comminucate with it.
func RegisterDiscovery(h bh.Hive) {
	a := h.NewApp("discovery")
	a.Handle(nom.NodeJoined{}, &nodeJoinedHandler{})
	a.Handle(nom.NodeLeft{}, &nodeLeftHandler{})

	a.Handle(nom.PortUpdated{}, &portUpdateHandler{})
	// TODO(soheil): Handle PortRemoved.

	a.Handle(nom.PacketIn{}, &lldpPktInHandler{})

	host := h.NewApp("discovery.Host")
	host.Handle(nom.PacketIn{}, &arpPktInHandler{})
	host.Handle(nom.HostConnected{}, &hostConnectedHandler{})

	a.Handle(NewLink{}, &newLinkHandler{})
	a.Handle(lldpTimeout{}, &timeoutHandler{})
	go func() {
		for {
			h.Emit(lldpTimeout{})
			time.Sleep(60 * time.Second)
		}
	}()
}
