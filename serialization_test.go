package gossip

import "testing"

import "gossip/members"
import "bytes"

func TestGossipMessage(t *testing.T) {

	gm := GossipMessage{
		To: members.GossipMember{
			ID: members.NewID(50, 51),
		},
		From: members.GossipMember{
			ID: members.NewID(50, 52),
		},
		Payload: []byte{1, 2, 3, 4, 5},
	}

	var buf bytes.Buffer

	gm.ToBytes(&buf)

	var gm2 GossipMessage

	gm2.FromBytes(&buf)

	if gm.To != gm2.To {
		t.Log("gm.To != gm2.To")
		t.Fail()
	}
	if gm.From != gm2.From {
		t.Log("gm.From != gm2.From")
		t.Fail()
	}
	if len(gm.Payload) == len(gm2.Payload) {
		for i := range gm.Payload {
			if gm.Payload[i] != gm2.Payload[i] {
				t.Log("gm.Payload[", i, "] != gm2.Payload[", i, "]")
				t.Fail()
			}
		}
	} else {
		t.Log("len(gm.Payload)", len(gm.Payload), " != gm2.Payload", len(gm2.Payload))
		t.Fail()
	}
}



func TestGossipSerialization(t *testing.T) {

	gm := GossipMessage{
		To: members.GossipMember{
			ID: members.NewID(50, 51),
		},
		From: members.GossipMember{
			ID: members.NewID(50, 52),
		},
		Payload: []byte{1, 2, 3, 4, 5},
	}

  g := Gossip {
    Type: DataMessage,
    Message: gm,
    Members: make([]members.GossipMember,10,10),
  }

  for i:=0; i< 10; i++ {
    g.Members[i] = members.GossipMember {ID: members.NewID(uint64(i), uint64(i)) }
  }

	var buf bytes.Buffer

	g.ToBytes(&buf)

  var g2 Gossip

  g2.FromBytes(&buf)

	if g.Type != g2.Type {
		t.Log("g.Type != g2.Type")
		t.Fail()
	}

	if len(g.Members) != len(g2.Members) {
		t.Log("len(g.Members)",len(g.Members)," != len(g2.Members", len(g2.Members))
		t.Fail()
	} else {
		for i := range g.Members {
			if g.Members[i] != g2.Members[i] {
				t.Log("g.Members[", i, "] =",g.Members[i])
				t.Log("g2.Members[", i, "] =",g2.Members[i])
				t.Log("g.Members[", i, "] != g2.Members[", i, "]")
				t.Fail()
			}
		}
  }

	gm2 := g2.Message

	if gm.To != gm2.To {
		t.Log("gm.To != gm2.To")
		t.Fail()
	}

	if gm.From != gm2.From {
		t.Log("gm.From != gm2.From")
		t.Fail()
	}

	if len(gm.Payload) == len(gm2.Payload) {
		for i := range gm.Payload {
			if gm.Payload[i] != gm2.Payload[i] {
				t.Log("gm.Payload[", i, "] != gm2.Payload[", i, "]")
				t.Fail()
			}
		}
	} else {
		t.Log("len(gm.Payload)", len(gm.Payload), " != gm2.Payload", len(gm2.Payload))
		t.Fail()
	}
}
