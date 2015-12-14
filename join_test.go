package gossip

import "gossip/members"
import "testing"

func TestRequestJoin(t *testing.T) {
	var a members.MemberAddress = members.MemberAddress{IP: 0x12345678, UDPPort: 0x4545, TCPPort: 0x5454}

	s := members.GossipMember{ID: members.NewID(8, 0)}

	res := RequestJoin(a, s)
	if res.Type != JoinRequest {
		t.Fail()
	}

	if res.Message.To.Address != a {
		t.Fail()
	}

	if res.Message.From != s {
		t.Fail()
	}
}
