/*
 * In The Name Of God
 * ========================================
 * [] File Name : arp.go
 *
 * [] Creation Date : 17-11-2015
 *
 * [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Elahe Jalalpour.
 */

package discovery

import (
	"errors"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	"github.com/elahejalalpour/beehive-netctrl/nom"
	//	bh "github.com/kandoo/beehive"
)

func decodeARP(b []byte) (nom.Host, nom.Port, error) {
	p := gopacket.NewPacket(b, layers.LayerTypeEthernet, gopacket.Default)
	arplayer := p.Layer(layers.LayerTypeARP)

	if arplayer == nil {
		return nom.Host{}, nom.Port{}, errors.New("decodeARP: no ARP layer in packet")
	}

	arp, _ := arplayer.(*layers.ARP)

	if arp.AddrType != layers.LinkTypeEthernet {
		return nom.Host{}, nom.Port{}, errors.New("decodeARP: layer 2 protocol is no supported")
	}

	h := nom.Host{}
	copy(h.MACAddr[:], arp.SourceHwAddress[:6])

	//portUID := nom.UID(v[1:])
	//hID, pID := nom.ParsePortUID(portUID)
	//h.ID = hID

	return h, nom.Port{}, nil
}
