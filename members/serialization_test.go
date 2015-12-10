package members

import "testing"
import "bytes"

func TestSerialization(t *testing.T) {
	var buf bytes.Buffer
	m := GossipMember{
		ID:        NewID(50, 50),
		heartbeat: 50,
		Address: MemberAddress{
			IP:      0x12345678,
			UDPPort: 0x5454,
			TCPPort: 0x4545,
		},
	}

	m.ToBytes(&buf)

	n := m

	err := m.FromBytes(&buf)
	if err != nil || n != m {
		t.Log("serialization error")
		t.Fail()
	}

}
