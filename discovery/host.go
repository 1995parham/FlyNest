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
	"encoding/json"
	"fmt"
	"github.com/elahejalalpour/beehive-netctrl/http"
	"github.com/elahejalalpour/beehive-netctrl/nom"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/golang/glog"
)

type httpHostListHandler struct{}

func (h *httpHostListHandler) Rcv(msg bh.Msg, ctx bh.RcvContext) error {
	hrq := msg.Data().(http.HTTPRequest)
	if hrq.AppName == "host" && hrq.Verb == "list" {
		dict := ctx.Dict(hostDict)

		v, err := dict.Get("hsts")
		hsts := []nom.Host{}
		if err == nil {
			hsts = v.([]nom.Host)
		}

		data, err := json.Marshal(hsts)
		if err != nil {
			glog.Errorf("Host list JSON marshaling: %v", err)
			return err
		}

		fmt.Println(hsts)

		hrs := http.HTTPResponse{
			AppName: "host",
			Data:    data,
		}

		err = ctx.Reply(msg, hrs)
		if err != nil {
			glog.Errorf("Replay error: %v", err)
			return err
		}
	}
	return nil
}

func (h *httpHostListHandler) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {
	return bh.MappedCells{
		{hostDict, "hsts"},
	}
}

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
		glog.Errorf("ARP decoding error: %v", err)
		return err
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
		return nil
	}

	v, err := dict.Get("hsts")
	hsts := []nom.Host{}
	if err == nil {
		hsts = v.([]nom.Host)

	}
	hsts = append(hsts, nom.Host(host))

	ctx.Emit(nom.HostJoined(host))

	err = dict.Put(host.MACAddr.String(), host)
	if err != nil {
		glog.Errorf("Put %v in %s: %v", host, host.MACAddr.String(), err)
		return err
	}

	err = dict.Put("hsts", hsts)
	if err != nil {
		glog.Errorf("Put %v in hsts: %v", hsts, err)
		return err
	}

	return nil
}

func (h *hostConnectedHandler) Map(msg bh.Msg, ctx bh.MapContext) bh.MappedCells {
	return bh.MappedCells{
		{hostDict, msg.Data().(nom.HostConnected).MACAddr.String()},
		{hostDict, "hsts"},
	}
}

const (
	hostDict = "H"
)
