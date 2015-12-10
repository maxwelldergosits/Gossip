package members
type MemberStatus int

//TODO figure out UUID stuff
type MemberID uint64

//TODO figure out address stuff
type MemberAddress int

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

func (m *GossipMember) ToBytes() []byte {
	return []byte{}
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


