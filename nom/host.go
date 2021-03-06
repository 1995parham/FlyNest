/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 24-11-2015
 * |
 * | File Name:     host.go
 * +===============================================
 */
package nom

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// TODO(elahe): add documentation ..
type HostLeft Host

// HostJoined is a message emitted when a host connects to network and
// we added it into our cells.
type HostJoined Host

// HostConnected is a message emitted when a host connects to network,
// we use this event in order to add host into our cells.
type HostConnected Host

// HostConnected is a message emitted when a host disconnects from network,
// we use this event in order to remove host from our cells.
type HostDisconnected Host

// Host represnts a end point element, such as your pc.
type Host struct {
	ID       HostID
	Net      UID
	MACAddr  MACAddr
	IPv4Addr IPv4Addr
	Node     UID
}

// HostID is the ID of a host. This must be unique among all hosts in
// the network.
type HostID string

func (h Host) String() string {
	return fmt.Sprintf("Host %s (mac=%v) (ip=%v)", string(h.ID), h.MACAddr, h.IPv4Addr)
}

// UID converts id into a UID.
func (id HostID) UID() UID {
	return UID(id)
}

// UID returns the node's unique ID. This id is in the form of net_id$$host_id.
func (n Host) UID() UID {
	return UID(string(n.ID))
}

// ParseHostUID parses a UID of a host and returns the respective host IDs.
func ParseHostUID(id UID) HostID {
	s := UIDSplit(id)
	return HostID(s[0])
}

// JSONDecode decodes the host from a byte array using JSON.
func (n *Host) JSONDecode(b []byte) error {
	return json.Unmarshal(b, n)
}

// JSONEncode encodes the host into a byte array using JSON.
func (n *Host) JSONEncode() ([]byte, error) {
	return json.Marshal(n)
}

func init() {
	gob.Register(HostConnected{})
	gob.Register(HostDisconnected{})
	gob.Register(HostLeft{})
	gob.Register(HostJoined{})
}
