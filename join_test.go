package gossip

import "gossip/members"
import "testing"

func TestRequestJoin(t *testing.T) {
	var a members.MemberAddress = members.MemberAddress(4)

	s := members.GossipMember{ID: 8}

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
