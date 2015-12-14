package members
import "encoding/binary"
import "fmt"

type MemberStatus int

type MemberID struct {
	Upper uint64
	Lower uint64
}

func (m * MemberID) ToString() string {
  array := make([]byte,16,16)
  binary.BigEndian.PutUint64(array[:8],m.Upper)
  binary.BigEndian.PutUint64(array[8:],m.Lower)
  return fmt.Sprintf("%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X",array[0],array[1],array[2],array[3],array[4],array[5],array[6],array[7],array[8],array[9],array[10],array[11],array[12],array[13],array[14],array[15])
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
	Status    MemberStatus
	Heartbeat MemberHeartbeat
	Lastheard MemberHeartbeat
}

type MemberHandler interface {
	GetMembers(uint) []GossipMember
	GetAllMembers() []GossipMember
	Add(GossipMember)
	Find(MemberID) (GossipMember, bool)
  MarkSuspected(round, suspectedThreshold MemberHeartbeat)
  DeleteExpired(round, suspectedThreshold MemberHeartbeat)
  NumberOfMembers() int
}

func UpdateMember(handler MemberHandler, member GossipMember, round MemberHeartbeat) {

	m, exists := handler.Find(member.ID)

	if exists {
		if member.Heartbeat > m.Heartbeat {
			// update from the member
			handler.Add(GossipMember{
				ID:        member.ID,
				Heartbeat: member.Heartbeat,
				Lastheard: round,
				Status:    MemberAlive,
				Address:   member.Address,
			})
		}
	} else {
		handler.Add(GossipMember{
			ID:        member.ID,
			Heartbeat: member.Heartbeat,
			Lastheard: round,
			Status:    MemberAlive,
			Address:   member.Address,
		})
	}
}
