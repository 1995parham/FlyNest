package openflow

import (
	"github.com/1995parham/flynest/nom"
	//"github.com/1995parham/flynest/openflow/of10"
	"fmt"
	"github.com/1995parham/flynest/openflow/of12"
)

func (d *of12Driver) handlePortStatus(rep of12.PortStatus, c *ofConn) error {
	rawp := rep.Desc()
	namep := rawp.Name()
	p := nom.Port{
		ID:      portNoToPortID(rawp.PortNo()),
		Name:    string(namep[:]),
		MACAddr: rawp.HwAddr(),
		Node:    c.NodeUID(),
	}
	if rep.Reason() == 0 {
		// Port Add :)
		fmt.Printf("%v Added\n", p)
		d.ofPorts[rawp.PortNo()] = &p
		d.nomPorts[p.UID()] = rawp.PortNo()
	} else if rep.Reason() == 1 {
		// Port Remove
		fmt.Printf("%v Removed\n", p)
		delete(d.ofPorts, rawp.PortNo())
		delete(d.nomPorts, p.UID())
	}

	nd := nom.Driver{
		BeeID: c.ctx.ID(),
		Role:  nom.DriverRoleDefault,
	}

	c.ctx.Emit(nom.PortStatusChanged{
		Port:   p,
		Driver: nd,
	})
	return nil
}
