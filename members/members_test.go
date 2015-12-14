package members

import "testing"

func TestUpdateMember(t *testing.T) {

	testHandler := CreateMemoryMemberHandler()

	m := GossipMember{
		ID:        NewID(50, 0),
		Heartbeat: 1,
		Lastheard: 1,
	}

	UpdateMember(&testHandler, m, 1)

	_, exists := testHandler.Find(m.ID)

	if !exists {
		t.Log("member with ID", m.ID, "wasn't found")
		t.Fail()
	}

	UpdateMember(&testHandler, m, 10)

	mem, exists := testHandler.Find(m.ID)

	if !exists || mem.Lastheard != 1 {
		t.Log("member with ID", m.ID, "wasn't found with correct Lastheard at, wanted", 1, "got", mem.Lastheard)
		t.Fail()
	}

	m.Heartbeat = 20
	UpdateMember(&testHandler, m, 10)

	mem, exists = testHandler.Find(m.ID)

	if !exists || mem.Lastheard != 10 {
		t.Log("member with ID", m.ID, "wasn't found with correct Lastheard at, wanted", 10, "got", mem.Lastheard)
		t.Fail()
	}

}

func TestNewID(t *testing.T) {

	m := NewID(1, 2)
	if m.upper != 1 || m.lower != 2 {
		t.Fail()
	}
}
