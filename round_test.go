package gossip

import "testing"
import "gossip/members"

type TestRounder struct {
	m *members.MemoryMemberHandler
	c GossipConf
	o chan Gossip
}

func (t *TestRounder) Conf() GossipConf {
	return t.c
}

func (t *TestRounder) MemberHandler() members.MemberHandler {
	return t.m
}
func (t *TestRounder) Outbound() chan<- Gossip {
	return t.o
}

func TestSendRoundMessage(t *testing.T) {

	h := members.CreateMemoryMemberHandler()
	for i := 0; i < 15; i++ {
		h.Add(members.GossipMember{ID: members.MemberID(i)})
	}

	ts := &TestRounder{m: &h, c: GossipConf{RoundSize: 10}, o: make(chan Gossip, 10)}

	SendRoundMessage(ts)

	var res Gossip
	select {
	case res = <-ts.o:
	default:
		t.Log("no message")
		t.Fail()
	}

	if _, exists := h.Find(res.Message.To.ID); !exists || res.Type != DataMessage || len(res.Members) != 10 {
		t.Log("res.Message.To is in the members:", exists, "wanted", true)
		t.Log("res.Type = ", res.Type, "wanted", DataMessage)
		t.Log("len(res.Members) = ", len(res.Members), "wanted", 10)
		t.Fail()
	}
}
