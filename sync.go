package gossip

func requestSync(cxt gossipContext) {

	members := cxt.MemberHandler().GetMembers(1)

  // ask members[0] for all of its info
	msg := Gossip{
		Type: SyncRequest,
		message: GossipMessage{
			To:   members[0],
			From: cxt.Conf().Self,
		},
	}

	cxt.Outbound() <- msg


  // send all of your info to members[0]
  sync(cxt, members[0])

}


// Send all members to m
func sync(cxt gossipContext, m GossipMember) {

	members := cxt.MemberHandler().GetAllMembers()

	msg := Gossip{
		Type: SyncRequest,
		message: GossipMessage{
			To:   m,
			From: cxt.Conf().Self,
		},
    members: members,
	}

	cxt.Outbound() <- msg

}
