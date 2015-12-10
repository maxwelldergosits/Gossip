package members

type MemberStatus int

type MemberID struct {
	upper uint64
	lower uint64
}

func NewID(up, lo uint64) MemberID {
	return MemberID{up, lo}
}

// FOR NOW IPV4 only
type MemberAddress struct {
	IP      uint32
	UDPPort uint16
	TCPPort uint16
}

type MemberHeartbeat uint64

const (
	MemberAlive     = iota
	MemberSuspected = iota
)

type GossipMember struct {
	ID        MemberID
	Address   MemberAddress
	status    MemberStatus
	heartbeat MemberHeartbeat
	lastheard MemberHeartbeat
}

type MemberHandler interface {
	GetMembers(uint) []GossipMember
	GetAllMembers() []GossipMember
	Add(GossipMember)
	Find(MemberID) (GossipMember, bool)
}

func UpdateMember(handler MemberHandler, member GossipMember, round MemberHeartbeat) {

	m, exists := handler.Find(member.ID)

	if exists {
		if member.heartbeat > m.heartbeat {
			// update from the member
			handler.Add(GossipMember{
				ID:        member.ID,
				heartbeat: member.heartbeat,
				lastheard: round,
				status:    MemberAlive,
				Address:   member.Address,
			})
		}
	} else {
		handler.Add(GossipMember{
			ID:        member.ID,
			heartbeat: member.heartbeat,
			lastheard: round,
			status:    MemberAlive,
			Address:   member.Address,
		})
	}
}
