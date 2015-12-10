package gossip

import "gossip/members"

type Syncer interface {
	MemberHandler() members.MemberHandler
	Conf() GossipConf
}

// Choose a random member to sync with
func RequestSync(s Syncer) Gossip {

	members := s.MemberHandler().GetMembers(1)

	// send all of your info to members[0]
	return Sync(s, members[0], SyncRequest)

}

// Send all members to m
func Sync(s Syncer, m members.GossipMember, t MessageType) Gossip {
	members := s.MemberHandler().GetAllMembers()
	return Gossip{
		Type: t,
		Message: GossipMessage{
			To:   m,
			From: s.Conf().Self,
		},
		Members: members,
	}
}
