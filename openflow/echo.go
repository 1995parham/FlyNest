package openflow

import (
	"github.com/kandoo/beehive-netctrl/openflow/of"
	"github.com/kandoo/beehive-netctrl/openflow/of10"
	"github.com/kandoo/beehive-netctrl/openflow/of12"
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/golang/glog"
)

func (d *of10Driver) handleEchoRequest(req of10.EchoRequest, c *ofConn) error {
	return doHandleEchoRequest(req.Header, of10.NewEchoReply().Header, c)
}

func (d *of12Driver) handleEchoRequest(req of12.EchoRequest, c *ofConn) error {
	return doHandleEchoRequest(req.Header, of12.NewEchoReply().Header, c)
}

func doHandleEchoRequest(req of.Header, res of.Header, c *ofConn) error {
	glog.V(2).Infof("Received echo request from %v", c.node)
	res.SetXid(req.Xid())
	err := c.WriteHeaders([]of.Header{res})
	if err != nil {
		return err
	}
	c.Flush()
	glog.V(2).Infof("Sent echo reply to %v", c.node)
	return nil
}
