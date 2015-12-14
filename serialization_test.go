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
