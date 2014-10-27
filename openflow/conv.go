package openflow

import (
	"strconv"

	"github.com/kandoo/beehive-netctrl/nom"
)

func datapathIDToNodeID(dpID uint64) nom.NodeID {
	return nom.NodeID(strconv.FormatUint(dpID, 16))
}

func nodeIDToDatapathID(nodeID nom.NodeID) (uint64, error) {
	return strconv.ParseUint(string(nodeID), 16, 64)
}

func portNoToPortID(portNo uint32) nom.PortID {
	return nom.PortID(strconv.FormatUint(uint64(portNo), 10))
}

func portIDToPortNo(portID nom.PortID) (uint32, error) {
	no, err := strconv.ParseUint(string(portID), 10, 32)
	return uint32(no), err
}

func datapathIDToMACAddr(dpID uint64) nom.MACAddr {
	var mac [6]byte
	for i := 0; i < len(mac); i++ {
		mac[i] = byte((dpID >> uint(len(mac)-1-i)) & 0xFF)
	}
	return mac
}
