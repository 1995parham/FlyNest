/*
 * In The Name Of God
 * ========================================
 * [] File Name : host.go
 *
 * [] Creation Date : 16-11-2015
 *
 * [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Elahe Jalalpour.
 */

package nom

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// TODO(elahe): add documentation ..
type HostLeft Host

/*
 * HostJoined is a message emitted when a host connects to network.
 */
type HostJoined Host

/*
 * Host represnts a end point element, such as pc :)
 */
type Host struct {
	ID       HostID
	Net      UID
	MACAddr  MACAddr
	IPv4Addr IPv4Addr
}

/*
 * HostID is the ID of a host. This must be unique among all hosts in the
 * network
 */
type HostID string

func (h Host) String() string {
	return fmt.Sprintf("Host %s (mac=%v)", string(h.ID), h.MACAddr)
}

/*
 * UID converts id into a UID.
 */
func (id HostID) UID() UID {
	return UID(id)
}

/*
 * UID returns the node's unique ID. This id is in the form of net_id$$host_id.
 */
func (n Host) UID() UID {
	return UID(string(n.ID))
}

/*
 * ParseHostUID parses a UID of a host and returns the respective host IDs.
 */
func ParseHostUID(id UID) HostID {
	s := UIDSplit(id)
	return HostID(s[0])
}

/*
 * JSONDecode decodes the host from a byte array using JSON.
 */
func (n *Host) JSONDecode(b []byte) error {
	return json.Unmarshal(b, n)
}

/*
 * JSONEncode encodes the host into a byte array using JSON.
 */
func (n *Host) JSONEncode() ([]byte, error) {
	return json.Marshal(n)
}

func init() {
	gob.Register(HostLeft{})
	gob.Register(HostJoined{})
}
