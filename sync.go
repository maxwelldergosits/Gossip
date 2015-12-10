package gossip

// Choose a random member to sync with
func RequestSync(cxt GossipContext) {

	members := cxt.MemberHandler().GetMembers(1)

	// send all of your info to members[0]
	Sync(cxt, members[0], SyncRequest)

}

// Send all members to m
func Sync(cxt GossipContext, m GossipMember, t MessageType) {

	members := cxt.MemberHandler().GetAllMembers()

	msg := Gossip{
		Type: t,
		Message: GossipMessage{
			To:   m,
			From: cxt.Conf().Self,
		},
		Members: members,
	}

	cxt.Outbound() <- msg

}
