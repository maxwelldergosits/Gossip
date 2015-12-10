package members

import "bytes"
import "encoding/binary"

func (m *GossipMember) ToBytes(b *bytes.Buffer) {
	binary.Write(b, binary.BigEndian, m.ID.upper)
	binary.Write(b, binary.BigEndian, m.ID.lower)
	binary.Write(b, binary.BigEndian, m.heartbeat)
	m.Address.ToBytes(b)
}

func (m *MemberAddress) ToBytes(b *bytes.Buffer) {
	binary.Write(b, binary.BigEndian, m.IP)
	binary.Write(b, binary.BigEndian, m.UDPPort)
	binary.Write(b, binary.BigEndian, m.TCPPort)
}

func (m *GossipMember) FromBytes(b *bytes.Buffer) error {
	binary.Read(b, binary.BigEndian, &m.ID.upper)
	binary.Read(b, binary.BigEndian, &m.ID.lower)
	binary.Read(b, binary.BigEndian, &m.heartbeat)
	return m.Address.FromBytes(b)
}

func (m *MemberAddress) FromBytes(b *bytes.Buffer) error {
	binary.Read(b, binary.BigEndian, &m.IP)
	binary.Read(b, binary.BigEndian, &m.UDPPort)
	err := binary.Read(b, binary.BigEndian, &m.TCPPort)
	return err
}
