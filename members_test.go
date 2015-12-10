package gossip

import "testing"

func TestAdd(t *testing.T) {

	testHandler := CreateMemoryMemberHandler()
	m := GossipMember{
		id:        50,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)

	mem_ret, exists := testHandler.Find(m.id)

	if !exists {
		t.Log("member with id", m.id, "wasn't returned")
		t.Fail()
	}

	if mem_ret != m {
		t.Log("wrong member returned")
		t.Fail()
	}

  for i:= 100; i < 10000; i++ {
    m.id = MemberID(i)
    testHandler.Add(m)

    mem_ret, exists := testHandler.Find(m.id)

    if !exists {
      t.Log("member with id", m.id, "wasn't returned")
      t.Fail()
    }

    if mem_ret != m {
      t.Log("wrong member returned")
      t.Fail()
    }
  }

}

func TestUpdate(t *testing.T) {
	testHandler := CreateMemoryMemberHandler()
	m := GossipMember{
		id:        50,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)

	mem_ret, exists := testHandler.Find(m.id)

	if !exists {
		t.Log("member with id", m.id, "wasn't returned")
		t.Fail()
	}

	if mem_ret != m {
		t.Log("wrong member returned")
		t.Fail()
	}

	m.heartbeat += 1
	m.lastheard += 1

	testHandler.Add(m)

	mem_ret, exists = testHandler.Find(m.id)

	if !exists {
		t.Log("updated member with id", m.id, "wasn't returned")
		t.Fail()
	}

	if mem_ret != m {
		t.Log("wrong updated member returned")
		t.Fail()
	}
}

func TestSuspect(t * testing.T) {

	testHandler := CreateMemoryMemberHandler()

	m := GossipMember{
		id:        50,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)

  testHandler.MarkSuspected(10, 4)

	mem_ret, exists := testHandler.Find(m.id)

	if !exists {
		t.Log("member with id", m.id, "wasn't returned")
		t.Fail()
	}

	if mem_ret.id != m.id {
		t.Log("wrong member returned")
		t.Fail()
	}

	if mem_ret.status != MemberSuspected {
		t.Log("wrong status returned")
		t.Fail()
	}

}

func TestDelete(t *  testing.T) {

	testHandler := CreateMemoryMemberHandler()

	m := GossipMember{
		id:        50,
		heartbeat: 1,
		lastheard: 1,
	}
	m2 := GossipMember{
		id:        51,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)
	testHandler.Add(m2)


  testHandler.DeleteMember(m.id)
  testHandler.DeleteMember(m2.id)


	_, exists := testHandler.Find(m.id)

	if exists {
		t.Log("member with id", m.id, "was returned")
		t.Fail()
	}

	_, exists = testHandler.Find(m2.id)

	if exists {
		t.Log("member with id", m2.id, "was returned")
		t.Fail()
	}
  t.Log("now testing opposite deletion order")
	testHandler = CreateMemoryMemberHandler()

	testHandler.Add(m)
	testHandler.Add(m)

  testHandler.DeleteMember(m2.id)
  testHandler.DeleteMember(m.id)

	_, exists = testHandler.Find(m.id)

	if exists {
		t.Log("member with id", m.id, "was returned")
		t.Fail()
	}

	_, exists = testHandler.Find(m2.id)

	if exists {
		t.Log("member with id", m2.id, "was returned")
		t.Fail()
	}

}
func TestExpire(t * testing.T) {

	testHandler := CreateMemoryMemberHandler()

	m := GossipMember{
		id:        50,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)

  testHandler.DeleteExpired(10, 4)

	_, exists := testHandler.Find(m.id)

	if exists {
		t.Log("member with id", m.id, "wasn't deleted")
		t.Fail()
	}
}

func TestUpdateMember (t * testing.T) {


	testHandler := CreateMemoryMemberHandler()

	m := GossipMember{
		id:        50,
		heartbeat: 1,
		lastheard: 1,
	}

  UpdateMember(&testHandler, m, 1)

	_, exists := testHandler.Find(m.id)

	if !exists {
		t.Log("member with id", m.id, "wasn't found")
		t.Fail()
	}

  UpdateMember(&testHandler, m, 10)

	mem, exists := testHandler.Find(m.id)

	if !exists || mem.lastheard != 1 {
		t.Log("member with id", m.id, "wasn't found with correct lastheard at, wanted",1,"got",mem.lastheard)
		t.Fail()
	}

  m.heartbeat = 20
  UpdateMember(&testHandler, m, 10)

	mem, exists = testHandler.Find(m.id)

	if !exists || mem.lastheard != 10 {
		t.Log("member with id", m.id, "wasn't found with correct lastheard at, wanted",10,"got",mem.lastheard)
		t.Fail()
	}

}
