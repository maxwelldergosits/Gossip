package gossip

import "gossip/members"

//import "log"

/*
  Generates a message containing cxt.Conf().RoundSize members.
  To be sent over a UDP connection and is assumed to be an unreliable message.
*/
type rounder interface {
	Conf() *GossipConf
	MemberHandler() members.MemberHandler
	Outbound() chan<- Gossip
}

func SendRoundMessage(r rounder) {
	//log.Println("gossip", r.Conf().Self.ID, "round:", r.Conf().Self.Heartbeat, "members:", r.MemberHandler().NumberOfMembers())
	members := r.MemberHandler().GetMembers(r.Conf().RoundSize + 1)
	if len(members) == 0 {
		return
	}
	//log.Println("sending to",members[0], "with msg", members[1:])
	r.Outbound() <- Gossip{
		Type: DataMessage,
		Message: GossipMessage{
			To:   members[0],
			From: r.Conf().Self,
		},
		Members: members[1:],
	}

}
