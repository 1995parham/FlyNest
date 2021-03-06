package openflow

import (
	"fmt"

	"github.com/1995parham/flynest/nom"
	"github.com/1995parham/flynest/openflow/of10"
	"github.com/1995parham/flynest/openflow/of12"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/golang/glog"
)

func (of *of10Driver) handlePacketIn(in of10.PacketIn, c *ofConn) error {
	inPort := in.InPort()
	// Ignore packet-ins on switch specific ports.
	if inPort > uint16(of10.PP_MAX) {
		glog.V(2).Infof("ignoring packet-in on %v", inPort)
		return nil
	}

	port, ok := of.ofPorts[inPort]
	if !ok {
		return fmt.Errorf("of10Driver: port not found %v", inPort)
	}

	glog.V(2).Infof("packet received: %v", in)

	nomIn := nom.PacketIn{
		Node:     c.node.UID(),
		InPort:   port.UID(),
		BufferID: nom.PacketBufferID(in.BufferId()),
	}
	nomIn.Packet = nom.Packet(in.Data())
	c.ctx.Emit(nomIn)

	return nil
}

func (of *of12Driver) handlePacketIn(in of12.PacketIn, c *ofConn) error {
	m := in.Match()
	if m.Type() == uint16(of12.PMT_STANDARD) {
		glog.Warningf("standard matches are not supported")
		return nil
	}

	var inPort uint32
	hasInPort := false

	xm, _ := of12.ToOXMatch(in.Match())
	for _, f := range xm.Fields() {
		if of12.IsOxmInPort(f) {
			xp, _ := of12.ToOxmInPort(f)
			inPort = xp.InPort()
			hasInPort = true
		}
	}

	if !hasInPort {
		glog.V(2).Infof("packet in does not have an input port")
		return nil
	}

	// Ignore packet-ins on switch specific ports.
	if inPort > uint32(of12.PP_MAX) {
		glog.V(2).Infof("ignoring packet-in on %v", inPort)
		return nil
	}

	port, ok := of.ofPorts[inPort]
	if !ok {
		return fmt.Errorf("of12Driver: port not found %v", inPort)
	}

	if glog.V(2) {
		glog.Infof("packet received: %v", in)
	}

	nomIn := nom.PacketIn{
		Node:     c.node.UID(),
		InPort:   port.UID(),
		BufferID: nom.PacketBufferID(in.BufferId()),
	}
	nomIn.Packet = nom.Packet(in.Data())
	c.ctx.Emit(nomIn)

	return nil
}
