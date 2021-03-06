package openflow

import (
	"errors"

	"github.com/1995parham/flynest/openflow/of10"
	"github.com/1995parham/flynest/openflow/of12"
)

func (of *of10Driver) handleFeaturesReply(rep of10.FeaturesReply, c *ofConn) error {
	return lateFeaturesReplyError()
}

func (of *of12Driver) handleFeaturesReply(rep of12.FeaturesReply, c *ofConn) error {
	return lateFeaturesReplyError()
}

func lateFeaturesReplyError() error {
	return errors.New("Cannot receive a features reply after handshake.")
}
