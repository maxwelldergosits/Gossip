package gossip

func round(cxt gossipContext) {

	members := cxt.MemberHandler().GetMembers(cxt.Conf().RoundSize + 1)

	msg := Gossip{
		to:      members[0],
		from:    cxt.Conf().Self,
		members: members[1:],
	}

	cxt.Outbound() <- msg

}

func requestSync(cxt gossipContext) {

}
