package gossip

import (
	"gossip/members"
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
	Members []members.GossipMember
}


