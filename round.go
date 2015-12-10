package gossip
import "gossip/members"

/*
  Generates a message containing cxt.Conf().RoundSize members.
  To be sent over a UDP connection and is assumed to be an unreliable message.
*/
type rounder interface {
  Conf() GossipConf
  MemberHandler() members.MemberHandler
}

func SendRoundMessage(r rounder) Gossip {
	members := r.MemberHandler().GetMembers(r.Conf().RoundSize + 1)

	return Gossip{
		Type: DataMessage,
		Message: GossipMessage{
			To:   members[0],
			From: r.Conf().Self,
		},
		Members: members[1:],
	}

}
