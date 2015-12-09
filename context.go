package gossip

import "time"

type GossipConf struct {
	RoundLength time.Duration
	SyncLength  time.Duration
	RoundSize   uint
	Self        GossipMember
}

func createContext(conf GossipConf, outbound, received chan GossipMessage) (gossipContext, error) {
	return &DefaultGossipContext{
		outbound:         make(chan Gossip, 1000),
		inbound:          make(chan Gossip, 1000),
		outboundMessages: outbound,
		receivedMessages: received,
		conf:             conf}, nil
}

type gossipContext interface {
	Inbound() chan Gossip
	Outbound() chan Gossip
	OutboundMessages() chan GossipMessage
	ReceivedMessages() chan GossipMessage
	MemberHandler() MemberHandler
	Conf() GossipConf
}

type DefaultGossipContext struct {
	inbound          chan Gossip
	outbound         chan Gossip
	outboundMessages chan GossipMessage
	receivedMessages chan GossipMessage
	memberHandler    MemberHandler
	conf             GossipConf
}

func (cxt *DefaultGossipContext) Inbound() chan Gossip {
	return cxt.inbound
}
func (cxt *DefaultGossipContext) Outbound() chan Gossip {
	return cxt.outbound
}

func (cxt *DefaultGossipContext) MemberHandler() MemberHandler {
	return cxt.memberHandler
}

func (cxt *DefaultGossipContext) OutboundMessages() chan GossipMessage {
	return cxt.outboundMessages
}

func (cxt *DefaultGossipContext) ReceivedMessages() chan GossipMessage {
	return cxt.receivedMessages
}

func (cxt *DefaultGossipContext) Conf() GossipConf {
	return cxt.conf
}
