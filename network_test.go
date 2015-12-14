package gossip

import (
	"gossip/members"
	"testing"
)

func TestUDPNetwork(t *testing.T) {

	addr := members.MemberAddress{
		IP:      [4]byte{0,0,0,0},
		UDPPort: 12000,
	}

	out := make(chan Gossip, 10)
	in := make(chan Gossip, 10)

	err := ListenUDP(addr.UDPPort, in)
	if err != nil {
		t.Log("error:", err.Error())
		t.Fail()
	}

	err = SendUDP(addr, out)
	if err != nil {
		t.Log("error:", err.Error())
		t.Fail()
	}

	g := Gossip{
		Message: GossipMessage{
			To: members.GossipMember{
				Address: members.MemberAddress{
					UDPPort: 12000,
		      IP:      [4]byte{0,0,0,0},
				},
			},
			Payload: []byte{1, 2, 3, 4},
		},
	}

	out <- g

	g2 := <-in

	t.Log(g2)

	if g2.Message.To != g.Message.To {
		t.Log("g2.To != g.To")
		t.Fail()
	}

	if len(g2.Message.Payload) != len(g.Message.Payload) {
		t.Log("len(g2.Message.Payload) != len(g.Message.Payload)")
		t.Fail()
	}

}
