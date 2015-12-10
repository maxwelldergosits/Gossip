package gossip

import "time"

type GossipConf struct {
	RoundLength time.Duration
	SyncLength  time.Duration
	RoundSize   uint
	Self        GossipMember
	Connect     bool
}

func createContext(conf GossipConf, outbound, received chan GossipMessage) (GossipContext, error) {
	return &DefaultGossipContext{
		outbound:         make(chan Gossip, 1000),
		inbound:          make(chan Gossip, 1000),
		outboundMessages: outbound,
		receivedMessages: received,
		conf:             conf,
		round:            0,
	}, nil
}

type GossipContext interface {
	Inbound() chan Gossip
	Outbound() chan Gossip
	OutboundMessages() chan GossipMessage
	ReceivedMessages() chan GossipMessage
	MemberHandler() MemberHandler
	Conf() GossipConf
	Round() MemberHeartbeat
}

type DefaultGossipContext struct {
	inbound          chan Gossip
	outbound         chan Gossip
	outboundMessages chan GossipMessage
	receivedMessages chan GossipMessage
	memberHandler    MemberHandler
	conf             GossipConf
	round            MemberHeartbeat
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

func (cxt *DefaultGossipContext) Round() MemberHeartbeat {
	return cxt.round
}
