package gossip
import "testing"
import "gossip/members"



func TestGossipContext(t * testing.T) {
  c := GossipConf{RoundSize:30}

  xm := make(chan GossipMessage)
  ym := make(chan GossipMessage)

  h := members.CreateMemoryMemberHandler()
  r := members.MemberHeartbeat(1)

  cxt, err := createContext(c, xm, ym)

  if err != nil {
    t.Log("err != nil", err.Error())
    t.Fail()
  }

  cxt.SetMemberHandler(&h)
  cxt.SetRound(r)


  if (cxt.Inbound() == cxt.Outbound()) {
    t.Fail()
  }
	if (cxt.OutboundMessages() != xm) {
    t.Fail()
  }
	if (cxt.ReceivedMessages() != ym) {
    t.Fail()
  }
	if (cxt.MemberHandler() != &h) {
    t.Fail()
  }
	if(cxt.Conf() != c) {
    t.Fail()
  }
	if(cxt.Round()!= r ){
    t.Fail()
  }

}

