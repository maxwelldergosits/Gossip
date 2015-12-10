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

func TestGetSyncRequest(t *testing.T) {

  conf := GossipConf {
    RoundLength:time.Duration(1000000000*time.Millisecond),
    SyncLength:time.Duration( 1000000000*time.Millisecond),
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

  msg := Gossip {
    Type:SyncRequest,
  }
  cxt.Inbound() <- msg

	go startGossip(cxt)

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

  if (msg.Type != DataMessage) {
    t.Log("msg.Type != SyncRequest, it is ", msg.Type)
    t.Fail()
  }

}

func TestGetJoinRequest(t *testing.T) {

  conf := GossipConf {
    RoundLength:time.Duration(1000000000*time.Millisecond),
    SyncLength:time.Duration( 1000000000*time.Millisecond),
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

  msg := Gossip {
    Type:JoinRequest,
  }
  cxt.Inbound() <- msg

	go startGossip(cxt)

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

  if (msg.Type != DataMessage) {
    t.Log("msg.Type != SyncRequest, it is ", msg.Type)
    t.Fail()
  }

}

func TestReceive(t *testing.T) {

  conf := GossipConf {
    RoundLength:time.Duration(1000000000*time.Millisecond),
    SyncLength:time.Duration( 1000000000*time.Millisecond),
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

  msg := Gossip {
    Type:DataMessage,
    Members:[]members.GossipMember{members.GossipMember{ID:members.MemberID(20000)}},
    Message:GossipMessage{From:members.GossipMember{ID:members.MemberID(20001)}},
  }
  cxt.Inbound() <- msg

	go startGossip(cxt)

  _ = <- time.Tick(2*time.Millisecond)

  _, exists := cxt.MemberHandler().Find(members.MemberID(20000))
  if !exists {
    t.Log("Member with ID 20000 not found")
    t.Fail()
  }
  _, exists  = cxt.MemberHandler().Find(members.MemberID(20001))
  if !exists {
    t.Log("Member with ID 20001 not found")
    t.Fail()
  }

}
