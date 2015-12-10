package gossip
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
	id        MemberID
	address   MemberAddress
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

func(m * MemoryMemberHandler)	GetMembers(uint) []GossipMember {
return []GossipMember{}
}

func(m * MemoryMemberHandler)	GetAllMembers() []GossipMember {

return []GossipMember{}
}

func UpdateMember(handler MemberHandler, member GossipMember, round MemberHeartbeat) {

	m, exists := handler.Find(member.id)

	if exists {
		if member.heartbeat > m.heartbeat {
			// update from the member
			handler.Add(GossipMember{
				id:        member.id,
				heartbeat: member.heartbeat,
				lastheard: round,
				status:    MemberAlive,
				address:   member.address,
			})
		}
	} else {
		handler.Add(GossipMember{
			id:        member.id,
			heartbeat: member.heartbeat,
			lastheard: round,
			status:    MemberAlive,
			address:   member.address,
		})
	}
}

type mmhEntry struct {
	member GossipMember
	index  uint
}

type MemoryMemberHandler struct {
	table map[MemberID]*mmhEntry
	list  []MemberID
}

func CreateMemoryMemberHandler() MemoryMemberHandler {
	return MemoryMemberHandler{
		table: make(map[MemberID]*mmhEntry),
		list:  make([]MemberID, 0, 1000),
	}
}

func (mmh *MemoryMemberHandler) Add(member GossipMember) {
  if entry, ok := mmh.table[member.id]; ok {
    entry.member = member
  } else {
    var m uint = uint(len(mmh.list))
    n := m + 1
    if n > uint(cap(mmh.list)) {
        newSlice := make([]MemberID, (n+1)*2)
        copy(newSlice, mmh.list)
        mmh.list = newSlice
    }
    mmh.list = mmh.list[0:n]
    mmh.list[m] = member.id
    mmh.table[member.id]= &mmhEntry{member, m}
  }
}

func (m *MemoryMemberHandler) Find(id MemberID) (GossipMember, bool) {
	entry, ok := m.table[id]
	if ok {
		return entry.member, ok
	} else {
		return GossipMember{}, false
	}
}

func (m * MemoryMemberHandler) MarkSuspected(round, suspectedThreshold MemberHeartbeat) {
  for _,val := range m.table {
    if (round - val.member.lastheard > suspectedThreshold) {
      val.member.status = MemberSuspected
    }
  }
}
func (m * MemoryMemberHandler) DeleteMember(id MemberID) {
  entry,ok := m.table[id]
  if !ok { return }
  if entry.index != uint(len(m.list)-1) {
    index := len(m.list)-1
    lastEntry := m.table[m.list[index]]
    lastEntry.index = entry.index
    m.list[lastEntry.index] = lastEntry.member.id
  }
  m.list = m.list[0:len(m.list)-1]
  delete(m.table, id)
}

func (m * MemoryMemberHandler) DeleteExpired(round, expireThreshold MemberHeartbeat) {
  dead := make([]MemberID, 1000)
  for _,val := range m.table {
    if (round - val.member.lastheard > expireThreshold) {
     dead = append(dead, val.member.id)
    }
  }
  for _, id := range dead {
    m.DeleteMember(id)
  }
}
