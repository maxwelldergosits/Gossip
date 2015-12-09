package gossip

import (
	"bytes"
)

type MessageType int

const (
	JoinRequest MessageType = iota
	DataMessage             = iota
	SyncRequest             = iota
)

type Gossip struct {
	Type    MessageType
	message GossipMessage
	members []GossipMember
}

func SerializeMessage(buffer bytes.Buffer, message Gossip) error {
	return nil
}

/* send a message to other hosts */
func sendMessage(cxt gossipContext, message GossipMessage) {

	cxt.Outbound() <- Gossip{
		Type:    DataMessage,
		message: message,
		members: []GossipMember{},
	}

}

/* internal message received from other gossip member */
func handleMessage(cxt gossipContext, gossip Gossip) {

	if gossip.message.To == cxt.Conf().Self {
		cxt.ReceivedMessages() <- gossip.message
	} else {
		//forwardMessage(cxt, message)
	}

	updateMember(cxt.MemberHandler(), gossip.message.From)
	for _, member := range gossip.members {
		updateMember(cxt.MemberHandler(), member)
	}

}
