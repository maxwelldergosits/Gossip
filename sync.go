package gossip

import "gossip/members"

type Syncer interface {
	MemberHandler() members.MemberHandler
	Conf() *GossipConf
	Outbound() chan<- Gossip
}

// Choose a random member to sync with
func RequestSync(s Syncer) {

	members := s.MemberHandler().GetMembers(1)

	if len(members) == 0 {
		return
	}
	// send all of your info to members[0]
	s.Outbound() <- Sync(s, members[0], SyncRequest)

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
