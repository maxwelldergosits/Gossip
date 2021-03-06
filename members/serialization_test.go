package members

import "testing"
import "bytes"

func TestSerialization(t *testing.T) {
	var buf bytes.Buffer
	m := GossipMember{
		ID:        NewID(50, 50),
			Address: MemberAddress {
      IP:      [4]byte{0x12,0x34,0x56,0x78},
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
