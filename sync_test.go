package gossip

import "gossip/members"
import "testing"

type TestSyncer struct {
	m *members.MemoryMemberHandler
	c GossipConf
}

func (t *TestSyncer) Conf() GossipConf {
	return t.c
}

func (t *TestSyncer) MemberHandler() members.MemberHandler {
	return t.m
}

func TestSync(t *testing.T) {

	h := members.CreateMemoryMemberHandler()

	h.Add(members.GossipMember{ID: 50})

	ts := &TestSyncer{m: &h}
	to := members.GossipMember{ID: 51}
	res := Sync(ts, to, DataMessage)

	if res.Message.To != to || res.Type != DataMessage {
		t.Log("res.Message.To = ", res.Message.To, "to =", to)
		t.Log("res.Type = ", res.Type)
		t.Fail()
	}

	if (len(res.Members) != 1 || res.Members[0] != members.GossipMember{ID: 50}) {
		t.Log("len(res.Members)= ", len(res.Members))
		t.Fail()
	}

}

func TestRequestSync(t *testing.T) {

	h := members.CreateMemoryMemberHandler()

	h.Add(members.GossipMember{ID: 50})

	ts := &TestSyncer{m: &h}
	res := RequestSync(ts)

	if _, exists := h.Find(res.Message.To.ID); !exists || res.Type != SyncRequest {
		t.Log("res.Type = ", res.Type)
		t.Fail()
	}

	if (len(res.Members) != 1 || res.Members[0] != members.GossipMember{ID: 50}) {
		t.Log("len(res.Members)= ", len(res.Members))
		t.Fail()
	}

}
