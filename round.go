package gossip
import "gossip/members"

/*
  Generates a message containing cxt.Conf().RoundSize members.
  To be sent over a UDP connection and is assumed to be an unreliable message.
*/
type rounder interface {
  Conf() GossipConf
  MemberHandler() members.MemberHandler
  Outbound() chan Gossip


}

func SendRoundMessage(r rounder) {
	members := r.MemberHandler().GetMembers(r.Conf().RoundSize + 1)
  if (len(members) == 0) {return}
	r.Outbound() <- Gossip{
		Type: DataMessage,
		Message: GossipMessage{
			To:   members[0],
			From: r.Conf().Self,
		},
		Members: members[1:],
	}

}
