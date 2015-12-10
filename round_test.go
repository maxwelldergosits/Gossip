package gossip
import "testing"
import "gossip/members"

type TestRounder struct {
  m *members.MemoryMemberHandler
  c GossipConf
}

func (t *TestRounder) Conf() GossipConf {
  return t.c
}

func (t *TestRounder) MemberHandler() members.MemberHandler {
  return t.m
}

func TestSendRoundMessage(t *testing.T) {

	h := members.CreateMemoryMemberHandler()
  for i:=0; i < 15; i++ {
	  h.Add(members.GossipMember{ID: members.MemberID(i)})
  }

	ts := &TestRounder{m: &h, c: GossipConf{RoundSize:10}}

	res := SendRoundMessage(ts)

	if _, exists := h.Find(res.Message.To.ID); !exists || res.Type != DataMessage || len(res.Members)!=10{
		t.Log("res.Message.To is in the members:",exists, "wanted", true)
		t.Log("res.Type = ", res.Type, "wanted",DataMessage)
		t.Log("len(res.Members) = ", len(res.Members), "wanted",10)
		t.Fail()
	}

}
