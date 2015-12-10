package members

import "testing"

func TestAdd(t *testing.T) {


	testHandler := CreateMemoryMemberHandler()
	m := GossipMember{
		ID:        50,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)

	mem_ret, exists := testHandler.Find(m.ID)

	if !exists {
		t.Log("member with ID", m.ID, "wasn't returned")
		t.Fail()
	}

	if mem_ret != m {
		t.Log("wrong member returned")
		t.Fail()
	}

  for i:= 100; i < 10000; i++ {
    m.ID = MemberID(i)
    testHandler.Add(m)

    mem_ret, exists := testHandler.Find(m.ID)

    if !exists {
      t.Log("member with ID", m.ID, "wasn't returned")
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
		ID:        50,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)

	mem_ret, exists := testHandler.Find(m.ID)

	if !exists {
		t.Log("member with ID", m.ID, "wasn't returned")
		t.Fail()
	}

	if mem_ret != m {
		t.Log("wrong member returned")
		t.Fail()
	}

	m.heartbeat += 1
	m.lastheard += 1

	testHandler.Add(m)

	mem_ret, exists = testHandler.Find(m.ID)

	if !exists {
		t.Log("updated member with ID", m.ID, "wasn't returned")
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
		ID:        50,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)

  testHandler.MarkSuspected(10, 4)

	mem_ret, exists := testHandler.Find(m.ID)

	if !exists {
		t.Log("member with ID", m.ID, "wasn't returned")
		t.Fail()
	}

	if mem_ret.ID != m.ID {
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
		ID:        50,
		heartbeat: 1,
		lastheard: 1,
	}
	m2 := GossipMember{
		ID:        51,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)
	testHandler.Add(m2)


  testHandler.DeleteMember(m.ID)
  testHandler.DeleteMember(m2.ID)


	_, exists := testHandler.Find(m.ID)

	if exists {
		t.Log("member with ID", m.ID, "was returned")
		t.Fail()
	}

	_, exists = testHandler.Find(m2.ID)

	if exists {
		t.Log("member with ID", m2.ID, "was returned")
		t.Fail()
	}
  t.Log("now testing opposite deletion order")
	testHandler = CreateMemoryMemberHandler()

	testHandler.Add(m)
	testHandler.Add(m)

  testHandler.DeleteMember(m2.ID)
  testHandler.DeleteMember(m.ID)

	_, exists = testHandler.Find(m.ID)

	if exists {
		t.Log("member with ID", m.ID, "was returned")
		t.Fail()
	}

	_, exists = testHandler.Find(m2.ID)

	if exists {
		t.Log("member with ID", m2.ID, "was returned")
		t.Fail()
	}

}
func TestExpire(t * testing.T) {

	testHandler := CreateMemoryMemberHandler()

	m := GossipMember{
		ID:        50,
		heartbeat: 1,
		lastheard: 1,
	}

	testHandler.Add(m)

  testHandler.DeleteExpired(10, 4)

	_, exists := testHandler.Find(m.ID)

	if exists {
		t.Log("member with ID", m.ID, "wasn't deleted")
		t.Fail()
	}
}

func TestGetN(t *testing.T) {

	testHandler := CreateMemoryMemberHandler()
  for i:= 0; i < 100; i++ {
    m := GossipMember{
      ID:        MemberID(i),
      heartbeat: 1,
      lastheard: 1,
    }
    testHandler.Add(m)
  }
  l := len(testHandler.GetMembers(10))
  if l != 10 {
    t.Log("didn't get enough random member, got",l, "wanted",10 )
    t.Fail()
  }

  if len(testHandler.GetMembers(1000)) != 100 {
    t.Log("got too many random members")
  }

	testHandler = CreateMemoryMemberHandler()

  l  = len(testHandler.GetMembers(10))

  if l != 0 {
    t.Log("GetMembers: got",l, "wanted",0 )
    t.Fail()
  }


}

func TestGetAll(t *testing.T) {

	testHandler := CreateMemoryMemberHandler()
  for i:= 0; i < 100; i++ {
    m := GossipMember{
      ID:        MemberID(i),
      heartbeat: 1,
      lastheard: 1,
    }
    testHandler.Add(m)
  }

  if len(testHandler.GetAllMembers()) != 100 {
    t.Log("didn't get enough random members")
  }

}
