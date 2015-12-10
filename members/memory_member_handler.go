package members

import "math/rand"

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
	if entry, ok := mmh.table[member.ID]; ok {
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
		mmh.list[m] = member.ID
		mmh.table[member.ID] = &mmhEntry{member, m}
	}
}

func (m *MemoryMemberHandler) Find(ID MemberID) (GossipMember, bool) {
	entry, ok := m.table[ID]
	if ok {
		return entry.member, ok
	} else {
		return GossipMember{}, false
	}
}

func (m *MemoryMemberHandler) MarkSuspected(round, suspectedThreshold MemberHeartbeat) {
	for _, val := range m.table {
		if round-val.member.lastheard > suspectedThreshold {
			val.member.status = MemberSuspected
		}
	}
}
func (m *MemoryMemberHandler) DeleteMember(ID MemberID) {
	entry, ok := m.table[ID]
	if !ok {
		return
	}
	if entry.index != uint(len(m.list)-1) {
		index := len(m.list) - 1
		lastEntry := m.table[m.list[index]]
		lastEntry.index = entry.index
		m.list[lastEntry.index] = lastEntry.member.ID
	}
	m.list = m.list[0 : len(m.list)-1]
	delete(m.table, ID)
}

func (m *MemoryMemberHandler) DeleteExpired(round, expireThreshold MemberHeartbeat) {
	dead := make([]MemberID, 1000)
	for _, val := range m.table {
		if round-val.member.lastheard > expireThreshold {
			dead = append(dead, val.member.ID)
		}
	}
	for _, ID := range dead {
		m.DeleteMember(ID)
	}
}

func (m *MemoryMemberHandler) GetMembers(n uint) []GossipMember {
	if n > uint(len(m.list)) {
		n = uint(len(m.list))
	}
	r := rand.New(rand.NewSource(99))
	members := make([]GossipMember, 0, n)
	var i uint = 0
	for ; i < n; i++ {
		r_index := r.Uint32() % uint32(len(m.list))
		member := m.table[m.list[r_index]].member
		members = append(members, member)
	}
	return members
}

func (m *MemoryMemberHandler) GetAllMembers() []GossipMember {

	members := make([]GossipMember, 0, 1000)
	for _, v := range m.table {
		members = append(members, v.member)
	}
	return members
}
