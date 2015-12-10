package gossip

import "gossip/members"

func RequestJoin(address members.MemberAddress, self members.GossipMember) Gossip {

	temp := members.GossipMember{
		Address: address,
	}
	return Gossip{
		Type: JoinRequest,
		Message: GossipMessage{
			To:   temp,
			From: self,
		},
	}

}
