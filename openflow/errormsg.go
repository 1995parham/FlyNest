package openflow

import (
	"github.com/kandoo/beehive/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/1995parham/flynest/openflow/of10"
	"github.com/1995parham/flynest/openflow/of12"
)

func (of *of10Driver) handleErrorMsg(err of10.ErrorMsg, c *ofConn) error {
	glog.Errorf("Error from switch %s: type=%d code=%d", c.node, err.ErrType(),
		err.Code())
	return nil
}

func (of *of12Driver) handleErrorMsg(err of12.ErrorMsg, c *ofConn) error {
	glog.Errorf("Error from switch %s: type=%d code=%d", c.node, err.ErrType(),
		err.Code())
	return nil
}
