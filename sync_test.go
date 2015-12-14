package gossip

import "gossip/members"
import "testing"

type TestSyncer struct {
	m *members.MemoryMemberHandler
	c GossipConf
	o chan Gossip
}

func (t *TestSyncer) Conf() *GossipConf {
	return &t.c
}

func (t *TestSyncer) MemberHandler() members.MemberHandler {
	return t.m
}
func (t *TestSyncer) Outbound() chan<- Gossip {
	return t.o
}

func TestSync(t *testing.T) {

	h := members.CreateMemoryMemberHandler()

	h.Add(members.GossipMember{ID: members.NewID(50, 0)})

	ts := &TestSyncer{m: &h}
	to := members.GossipMember{ID: members.NewID(51, 0)}

	res := Sync(ts, to, DataMessage)

	if res.Message.To != to || res.Type != DataMessage {
		t.Log("res.Message.To = ", res.Message.To, "to =", to)
		t.Log("res.Type = ", res.Type)
		t.Fail()
	}

	if (len(res.Members) != 1 || res.Members[0] != members.GossipMember{ID: members.NewID(50, 0)}) {
		t.Log("len(res.Members)= ", len(res.Members))
		t.Fail()
	}

}

func TestRequestSync(t *testing.T) {

	h := members.CreateMemoryMemberHandler()

	ts := &TestSyncer{m: &h, o: make(chan Gossip, 1000)}

	RequestSync(ts)

	var res Gossip

	select {
	case res = <-ts.o:
		t.Log("got message")
		t.Fail()
	default:
	}

	h.Add(members.GossipMember{ID: members.NewID(50, 0)})

	RequestSync(ts)

	select {
	case res = <-ts.o:
	default:
		t.Log("no message")
		t.Fail()
	}

	if _, exists := h.Find(res.Message.To.ID); !exists || res.Type != SyncRequest {
		t.Log("res.Type = ", res.Type)
		t.Fail()
	}

	if (len(res.Members) != 1 || res.Members[0] != members.GossipMember{ID: members.NewID(50, 0)}) {
		t.Log("len(res.Members)= ", len(res.Members))
		t.Fail()
	}

}
