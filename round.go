package gossip

func round(cxt gossipContext) {

	members := cxt.MemberHandler().GetMembers(cxt.Conf().RoundSize + 1)

	msg := Gossip{
		Type: DataMessage,
		message: GossipMessage{
			To:   members[0],
			From: cxt.Conf().Self,
		},
		members: members[1:],
	}

	cxt.Outbound() <- msg

}

