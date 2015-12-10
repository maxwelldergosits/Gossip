package gossip

import "testing"
import "gossip/members"
import "time"

type TestContext int

func TestRound(t *testing.T) {

  conf := GossipConf {
    RoundLength:time.Duration(100*time.Millisecond),
    SyncLength:time.Duration(1000000*time.Millisecond),
    RoundSize:20,
  }

  rec := make(chan GossipMessage, 1000)
  out := make(chan GossipMessage, 1000)
	cxt, _ := createContext(conf, rec, out)
  h := members.CreateMemoryMemberHandler()
  for i:=0; i < 1000; i++{
    h.Add(members.GossipMember{ID:members.MemberID(i)})
  }
  cxt.SetMemberHandler(&h)
	go startGossip(cxt)

  var msg Gossip
  gotMessage := false
  select {
    case msg = <- cxt.Outbound():
      gotMessage = true
    case _ = <- time.Tick(conf.RoundLength*2):
      gotMessage = false
  }
  if (!gotMessage) {
    t.Log("Never got a message")
    t.Fail()
  }
  if (len(msg.Members) != 20) {
    t.Log("len(msg.Members) wrong, got",len(msg.Members), "wanted 20")
    t.Fail()
  }

}

func TestSyncGossip(t *testing.T) {

  conf := GossipConf {
    RoundLength:time.Duration(1000000000*time.Millisecond),
    SyncLength:time.Duration(100*time.Millisecond),
    RoundSize:20,
  }

  rec := make(chan GossipMessage, 1000)
  out := make(chan GossipMessage, 1000)
	cxt, _ := createContext(conf, rec, out)
  h := members.CreateMemoryMemberHandler()
  for i:=0; i < 1000; i++{
    h.Add(members.GossipMember{ID:members.MemberID(i)})
  }
  cxt.SetMemberHandler(&h)
	go startGossip(cxt)

  var msg Gossip
  gotMessage := false

  select {
    case msg = <- cxt.Outbound():
      gotMessage = true
    case _ = <- time.Tick(conf.RoundLength*2):
      gotMessage = false
  }

  if (!gotMessage) {
    t.Log("Never got a message")
    t.Fail()
  }

  if (len(msg.Members) != 1000) {
    t.Log("len(msg.Members) wrong, got",len(msg.Members), "wanted 1000")
    t.Fail()
  }

  if (msg.Type != SyncRequest) {
    t.Log("msg.Type != SyncRequest, it is ", msg.Type)
    t.Fail()
  }

}
