/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 24-11-2015
 * |
 * | File Name:     arp.go
 * +===============================================
 */
package discovery

import (
	"errors"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	"github.com/1995parham/flynest/nom"
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
	copy(h.IPv4Addr[:], arp.SourceProtAddress[:4])

	h.ID = nom.HostID(h.MACAddr.String())

	return h, nom.Port{}, nil
}
