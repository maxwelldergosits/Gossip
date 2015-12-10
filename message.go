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
	Message GossipMessage
	Members []GossipMember
}

func SerializeMessage(buffer bytes.Buffer, message Gossip) error {
	return nil
}

/* send a message to other hosts */
func SendMessage(cxt GossipContext, message GossipMessage) {

	cxt.Outbound() <- Gossip{
		Type:    DataMessage,
		Message: message,
		Members: []GossipMember{},
	}

}

/* internal message received from other gossip member */
func HandleMessage(cxt GossipContext, gossip Gossip) {

	if gossip.Type == SyncRequest {
		Sync(cxt, gossip.Message.From, DataMessage)
	}

	if gossip.Type == JoinRequest {
		HandleJoin(cxt, gossip)
	}

	if gossip.Message.To == cxt.Conf().Self {
		cxt.ReceivedMessages() <- gossip.Message
	} else {
		//forwardMessage(cxt, message)
	}

	UpdateMember(cxt.MemberHandler(), gossip.Message.From, cxt.Round())
	for _, member := range gossip.Members {
		UpdateMember(cxt.MemberHandler(), member, cxt.Round())
	}

}
