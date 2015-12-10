package gossip

func RequestJoin(address MemberAddress, self GossipMember) Gossip {

	temp := GossipMember{
		address: address,
	}
	return Gossip{
		Type: JoinRequest,
		Message: GossipMessage{
			To:   temp,
			From: self,
		},
	}

}

func HandleJoin(cxt GossipContext, g Gossip) {

	cxt.Outbound() <- Gossip{
		Type: DataMessage,
		Message: GossipMessage{
			To:   g.Message.From,
			From: cxt.Conf().Self,
		},
		Members: cxt.MemberHandler().GetAllMembers(),
	}
}
