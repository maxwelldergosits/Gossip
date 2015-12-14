package gossip

import "testing"
import "gossip/members"
import "time"

type TestContext int

func TestRound(t *testing.T) {

	conf := GossipConf{
		RoundLength: time.Duration(100 * time.Millisecond),
		SyncLength:  time.Duration(1000000 * time.Millisecond),
		RoundSize:   20,
		Self:        members.GossipMember{ID: members.NewID(40000, 0)},
	}

	rec := make(chan GossipMessage, 1000)
	out := make(chan GossipMessage, 1000)
	cxt, _ := CreateContext(conf, rec, out)
	h := members.CreateMemoryMemberHandler()
	for i := 0; i < 1000; i++ {
		h.Add(members.GossipMember{ID: members.NewID(uint64(i), 0)})
	}
	cxt.SetMemberHandler(&h)

	go RunGossip(cxt)

	var msg Gossip
	gotMessage := false
	select {
	case msg = <-cxt.OutboundChannel:
		gotMessage = true
	case _ = <-time.Tick(conf.RoundLength * 2):
		gotMessage = false
	}
	if !gotMessage {
		t.Log("Never got a message")
		t.Fail()
	}
	if len(msg.Members) != 20 {
		t.Log("len(msg.Members) wrong, got", len(msg.Members), "wanted 20")
		t.Fail()
	}
	if msg.Message.From.Heartbeat != 1 {
		t.Log("wrong hb")
		t.Fail()
	}
}

func TestEmptyRound(t *testing.T) {

	conf := GossipConf{
		RoundLength: time.Duration(100 * time.Millisecond),
		SyncLength:  time.Duration(1000000 * time.Millisecond),
		RoundSize:   20,
		Self:        members.GossipMember{ID: members.NewID(40000, 0)},
	}

	rec := make(chan GossipMessage, 1000)
	out := make(chan GossipMessage, 1000)
	cxt, _ := CreateContext(conf, rec, out)
	h := members.CreateMemoryMemberHandler()
	cxt.SetMemberHandler(&h)
	go RunGossip(cxt)

	gotMessage := false
	select {
	case _ = <-cxt.OutboundChannel:
		gotMessage = true
	case _ = <-time.Tick(conf.RoundLength * 2):
		gotMessage = false
	}
	if gotMessage {
		t.Log("got a message")
		t.Fail()
	}

}

func TestSyncGossip(t *testing.T) {

	conf := GossipConf{
		RoundLength: time.Duration(1000000000 * time.Millisecond),
		SyncLength:  time.Duration(100 * time.Millisecond),
		RoundSize:   20,
		Self:        members.GossipMember{ID: members.NewID(40000, 0)},
	}

	rec := make(chan GossipMessage, 1000)
	out := make(chan GossipMessage, 1000)
	cxt, _ := CreateContext(conf, rec, out)
	h := members.CreateMemoryMemberHandler()
	for i := 0; i < 1000; i++ {
		h.Add(members.GossipMember{ID: members.NewID(uint64(i), 0)})
	}
	cxt.SetMemberHandler(&h)
	go RunGossip(cxt)

	var msg Gossip
	gotMessage := false

	select {
	case msg = <-cxt.OutboundChannel:
		gotMessage = true
	case _ = <-time.Tick(conf.RoundLength * 2):
		gotMessage = false
	}

	if !gotMessage {
		t.Log("Never got a message")
		t.Fail()
	}

	if len(msg.Members) != 1000 {
		t.Log("len(msg.Members) wrong, got", len(msg.Members), "wanted 1000")
		t.Fail()
	}

	if msg.Type != SyncRequest {
		t.Log("msg.Type != SyncRequest, it is ", msg.Type)
		t.Fail()
	}

}

func TestGetSyncRequest(t *testing.T) {

	conf := GossipConf{
		RoundLength: time.Duration(1000000000 * time.Millisecond),
		SyncLength:  time.Duration(1000000000 * time.Millisecond),
		RoundSize:   20,
		Self:        members.GossipMember{ID: members.NewID(40000, 0)},
	}

	rec := make(chan GossipMessage, 1000)
	out := make(chan GossipMessage, 1000)

	cxt, _ := CreateContext(conf, rec, out)
	h := members.CreateMemoryMemberHandler()
	for i := 0; i < 1000; i++ {
		h.Add(members.GossipMember{ID: members.NewID(uint64(i), 0)})
	}
	cxt.SetMemberHandler(&h)

	msg := Gossip{
		Type: SyncRequest,
	}
	cxt.InboundChannel <- msg

	go RunGossip(cxt)

	gotMessage := false

	select {
	case msg = <-cxt.OutboundChannel:
		gotMessage = true
	case _ = <-time.Tick(conf.RoundLength * 2):
		gotMessage = false
	}

	if !gotMessage {
		t.Log("Never got a message")
		t.Fail()
	}

	if len(msg.Members) != 1000 {
		t.Log("len(msg.Members) wrong, got", len(msg.Members), "wanted 1000")
		t.Fail()
	}

	if msg.Type != DataMessage {
		t.Log("msg.Type != SyncRequest, it is ", msg.Type)
		t.Fail()
	}

}

func TestGetJoinRequest(t *testing.T) {

	conf := GossipConf{
		RoundLength: time.Duration(1000000000 * time.Millisecond),
		SyncLength:  time.Duration(1000000000 * time.Millisecond),
		RoundSize:   20,
		Self:        members.GossipMember{ID: members.NewID(40000, 0)},
	}

	rec := make(chan GossipMessage, 1000)
	out := make(chan GossipMessage, 1000)

	cxt, _ := CreateContext(conf, rec, out)
	h := members.CreateMemoryMemberHandler()
	for i := 0; i < 1000; i++ {
		h.Add(members.GossipMember{ID: members.NewID(uint64(i), 0)})
	}
	cxt.SetMemberHandler(&h)

	msg := Gossip{
		Type: JoinRequest,
	}
	cxt.InboundChannel <- msg

	go RunGossip(cxt)

	gotMessage := false

	select {
	case msg = <-cxt.OutboundChannel:
		gotMessage = true
	case _ = <-time.Tick(conf.RoundLength * 2):
		gotMessage = false
	}

	if !gotMessage {
		t.Log("Never got a message")
		t.Fail()
	}

	if len(msg.Members) != 1000 {
		t.Log("len(msg.Members) wrong, got", len(msg.Members), "wanted 1000")
		t.Fail()
	}

	if msg.Type != DataMessage {
		t.Log("msg.Type != SyncRequest, it is ", msg.Type)
		t.Fail()
	}

}

func TestReceive(t *testing.T) {

	conf := GossipConf{
		RoundLength: time.Duration(1000000000 * time.Millisecond),
		SyncLength:  time.Duration(1000000000 * time.Millisecond),
		RoundSize:   20,
		Self:        members.GossipMember{ID: members.NewID(40000, 0)},
	}

	rec := make(chan GossipMessage, 1000)
	out := make(chan GossipMessage, 1000)

	cxt, _ := CreateContext(conf, rec, out)
	h := members.CreateMemoryMemberHandler()
	for i := 0; i < 1000; i++ {
		h.Add(members.GossipMember{ID: members.NewID(uint64(i), 0)})
	}
	cxt.SetMemberHandler(&h)

	msg := Gossip{
		Type:    DataMessage,
		Members: []members.GossipMember{members.GossipMember{ID: members.NewID(20000, 0)}},
		Message: GossipMessage{From: members.GossipMember{ID: members.NewID(20001, 0)}},
	}
	cxt.InboundChannel <- msg

	go RunGossip(cxt)

	_ = <-time.Tick(2 * time.Millisecond)

	_, exists := cxt.MemberHandler().Find(members.NewID(20000, 0))
	if !exists {
		t.Log("Member with ID 20000 not found")
		t.Fail()
	}
	_, exists = cxt.MemberHandler().Find(members.NewID(20001, 0))
	if !exists {
		t.Log("Member with ID 20001 not found")
		t.Fail()
	}

}

func TestReceiveMessage(t *testing.T) {

	conf := GossipConf{
		RoundLength: time.Duration(1000000000 * time.Millisecond),
		SyncLength:  time.Duration(1000000000 * time.Millisecond),
		RoundSize:   20,
		Self:        members.GossipMember{ID: members.NewID(40000, 0)},
	}

	rec := make(chan GossipMessage, 1000)
	out := make(chan GossipMessage, 1000)

	cxt, _ := CreateContext(conf, out, rec)
	h := members.CreateMemoryMemberHandler()
	for i := 0; i < 1000; i++ {
		h.Add(members.GossipMember{ID: members.NewID(uint64(i), 0)})
	}
	cxt.SetMemberHandler(&h)

	msg := Gossip{
		Type: DataMessage,
		Message: GossipMessage{
			From:    members.GossipMember{ID: members.NewID(20001, 0)},
			To:      members.GossipMember{ID: members.NewID(40000, 0)},
			Payload: []byte{1, 2},
		},
	}
	cxt.InboundChannel <- msg

	go RunGossip(cxt)

	// wait 2 millis
	var newMsg GossipMessage
	select {
	case newMsg = <-rec:
		if (newMsg.From != members.GossipMember{ID: members.NewID(20001, 0)}) || (newMsg.To != members.GossipMember{ID: members.NewID(40000, 0)}) {
			t.Log("got wrong message!")
			t.Fail()
		}
		if (len(newMsg.Payload) != 2) || newMsg.Payload[0] != 1 || newMsg.Payload[1] != 2 {
			t.Log("got wrong payload!")
			t.Fail()
		}
	case _ = <-time.Tick(2 * time.Millisecond):
		t.Log("test timed out!")
		t.Fail()
	}

}
func TestSendMessage(t *testing.T) {

	conf := GossipConf{
		RoundLength: time.Duration(1000000000 * time.Millisecond),
		SyncLength:  time.Duration(1000000000 * time.Millisecond),
		RoundSize:   20,
		Self:        members.GossipMember{ID: members.NewID(40000, 0)},
	}

	rec := make(chan GossipMessage, 1000)
	out := make(chan GossipMessage, 1000)

	cxt, _ := CreateContext(conf, out, rec)
	h := members.CreateMemoryMemberHandler()
	for i := 0; i < 1000; i++ {
		h.Add(members.GossipMember{ID: members.NewID(uint64(i), 0)})
	}
	cxt.SetMemberHandler(&h)

	out <- GossipMessage{
		From:    members.GossipMember{ID: members.NewID(40000, 0)},
		To:      members.GossipMember{ID: members.NewID(20001, 0)},
		Payload: []byte{1, 2},
	}

	go RunGossip(cxt)

	// wait 2 millis
	var newMsg Gossip
	select {
	case newMsg = <-cxt.OutboundChannel:
		if (newMsg.Message.From != members.GossipMember{ID: members.NewID(40000, 0)}) || (newMsg.Message.To != members.GossipMember{ID: members.NewID(20001, 0)}) {
			t.Log("got wrong message!")
			t.Fail()
		}
		if (len(newMsg.Message.Payload) != 2) || newMsg.Message.Payload[0] != 1 || newMsg.Message.Payload[1] != 2 {
			t.Log("got wrong payload!")
			t.Fail()
		}
		if newMsg.Type != DataMessage {
			t.Log("wrong message type")
			t.Fail()
		}
		if len(newMsg.Members) != 0 {
			t.Log("send members when it wasn't supposed to")
			t.Fail()
		}
	case _ = <-time.Tick(2 * time.Millisecond):
		t.Log("test timed out!")
		t.Fail()
	}

}
