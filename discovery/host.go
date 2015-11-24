/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 24-11-2015
 * |
 * | File Name:     host.go
 * +===============================================
 */
package discovery

import (
	"fmt"
	"github.com/elahejalalpour/beehive-netctrl/nom"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/golang/glog"
)

type arpPktInHandler struct{}

func (h *arpPktInHandler) Rcv(msg bh.Msg, ctx bh.RcvContext) error {
	pin := msg.Data().(nom.PacketIn)
	p := gopacket.NewPacket([]byte(pin.Packet), layers.LayerTypeEthernet, gopacket.Default)
	etherlayer := p.Layer(layers.LayerTypeEthernet)

	if etherlayer == nil {
		return nil
	}
	e, _ := etherlayer.(*layers.Ethernet)

	if e.EthernetType != layers.EthernetTypeARP {
		return nil
	}

	host, _, err := decodeARP([]byte(pin.Packet))

	if err != nil {
		return fmt.Errorf("ARP decoding error: %v", err)
	}
	glog.V(2).Infof("Host detected: %v", host)

	ctx.Emit(nom.HostConnected(host))

	return nil
}

func (h *arpPktInHandler) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {
	return bh.MappedCells{}
}

type hostConnectedHandler struct{}

func (h *hostConnectedHandler) Rcv(msg bh.Msg, ctx bh.RcvContext) error {
	host := msg.Data().(nom.HostConnected)
	dict := ctx.Dict(hostDict)

	_, err := dict.Get(host.MACAddr.String())
	if err == nil {
		glog.Warningf("Host rejoins: %v", host)
	}

	ctx.Emit(nom.HostJoined(host))

	return dict.Put(host.MACAddr.String(), host)
}

func (h *hostConnectedHandler) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {
	return bh.MappedCells{{hostDict, msg.Data().(nom.HostConnected).MACAddr.String()}}
}

const (
	hostDict = "H"
)
