package gossip

type GossipMember struct {
}

type MemberHandler interface {
	GetMembers(uint) []GossipMember
	GetAllMembers() []GossipMember
}

type MemoryMemberHandler struct {
}

func (m *GossipMember) ToBytes() []byte {

	return []byte{}

}

func updateMember(handler MemberHandler, member GossipMember) {

}

func (m *MemoryMemberHandler) GetMembers(n uint) []GossipMember {

	return []GossipMember{}

}
