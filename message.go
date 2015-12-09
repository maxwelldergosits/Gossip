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
	from    GossipMember
	to      GossipMember
	members []GossipMember
	payload GossipMessage
}

func SerializeMessage(buffer bytes.Buffer, message Gossip) error {
	buffer.Write(message.to.ToBytes())
	buffer.Write(message.from.ToBytes())
	return nil
}

/* send a message to other hosts */
func sendMessage(cxt gossipContext, message GossipMessage) {

}

/* internal message received from other gossip member */
func handleMessage(cxt gossipContext, message Gossip) {

	if message.to == cxt.Conf().Self {
		cxt.ReceivedMessages() <- message.payload
	} else {
		//forwardMessage(cxt, message)
	}
	updateMember(cxt.MemberHandler(), message.from)
	for _, member := range message.members {
		updateMember(cxt.MemberHandler(), member)
	}

}
