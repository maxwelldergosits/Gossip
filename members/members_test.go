package members

import "testing"


func TestUpdateMember (t * testing.T) {


	testHandler := CreateMemoryMemberHandler()

	m := GossipMember{
		ID:        50,
		heartbeat: 1,
		lastheard: 1,
	}

  UpdateMember(&testHandler, m, 1)

	_, exists := testHandler.Find(m.ID)

	if !exists {
		t.Log("member with ID", m.ID, "wasn't found")
		t.Fail()
	}

  UpdateMember(&testHandler, m, 10)

	mem, exists := testHandler.Find(m.ID)

	if !exists || mem.lastheard != 1 {
		t.Log("member with ID", m.ID, "wasn't found with correct lastheard at, wanted",1,"got",mem.lastheard)
		t.Fail()
	}

  m.heartbeat = 20
  UpdateMember(&testHandler, m, 10)

	mem, exists = testHandler.Find(m.ID)

	if !exists || mem.lastheard != 10 {
		t.Log("member with ID", m.ID, "wasn't found with correct lastheard at, wanted",10,"got",mem.lastheard)
		t.Fail()
	}

}
