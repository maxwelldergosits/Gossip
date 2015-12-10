package gossip

/*
  Generates a message containing cxt.Conf().RoundSize members.
  To be sent over a UDP connection and is assumed to be an unreliable message.
*/
func SendRoundMessage(cxt GossipContext) {

	members := cxt.MemberHandler().GetMembers(cxt.Conf().RoundSize + 1)

	msg := Gossip{
		Type: DataMessage,
		Message: GossipMessage{
			To:   members[0],
			From: cxt.Conf().Self,
		},
		Members: members[1:],
	}

	cxt.Outbound() <- msg

}
