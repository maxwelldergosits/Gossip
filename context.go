package gossip

import "gossip/members"
import "time"

type GossipConf struct {
	RoundLength time.Duration
	SyncLength  time.Duration
	RoundSize   uint
	Self        members.GossipMember
	Connect     bool
}

func createContext(conf GossipConf, outbound, received chan GossipMessage) (*DefaultGossipContext, error) {
	return &DefaultGossipContext{
		outbound:         make(chan Gossip, 1000),
		inbound:          make(chan Gossip, 1000),
		outboundMessages: outbound,
		receivedMessages: received,
		conf:             conf,
	}, nil
}

type GossipContext interface {
	Inbound() <-chan Gossip
	Outbound() chan<- Gossip
	OutboundMessages() chan GossipMessage
	ReceivedMessages() chan GossipMessage
	Conf() GossipConf
	MemberHandler() members.MemberHandler
	SetMemberHandler(members.MemberHandler)
	SetRound(members.MemberHeartbeat)
	Round() members.MemberHeartbeat
}

type DefaultGossipContext struct {
	inbound          chan Gossip
	outbound         chan Gossip
	outboundMessages chan GossipMessage
	receivedMessages chan GossipMessage
	memberHandler    members.MemberHandler
	conf             GossipConf
	round            members.MemberHeartbeat
}

func (cxt *DefaultGossipContext) Inbound() <-chan Gossip {
	return cxt.inbound
}
func (cxt *DefaultGossipContext) Outbound() chan<- Gossip {
	return cxt.outbound
}

func (cxt *DefaultGossipContext) MemberHandler() members.MemberHandler {
	return cxt.memberHandler
}

func (cxt *DefaultGossipContext) SetMemberHandler(m members.MemberHandler) {
	cxt.memberHandler = m
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

func (cxt *DefaultGossipContext) Round() members.MemberHeartbeat {
	return cxt.round
}

func (cxt *DefaultGossipContext) SetRound(r members.MemberHeartbeat) {
	cxt.round = r
}
