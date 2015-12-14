package members

import "bytes"
import "encoding/binary"

func (m *GossipMember) ToBytes(b *bytes.Buffer) {
	binary.Write(b, binary.BigEndian, m.ID.Upper)
	binary.Write(b, binary.BigEndian, m.ID.Lower)
	binary.Write(b, binary.BigEndian, m.Heartbeat)
	m.Address.ToBytes(b)
}

func (m *MemberAddress) ToBytes(b *bytes.Buffer) {
	binary.Write(b, binary.BigEndian, m.IP)
	binary.Write(b, binary.BigEndian, m.UDPPort)
	binary.Write(b, binary.BigEndian, m.TCPPort)
}

func (m *GossipMember) FromBytes(b *bytes.Buffer) error {
	binary.Read(b, binary.BigEndian, &m.ID.Upper)
	binary.Read(b, binary.BigEndian, &m.ID.Lower)
	binary.Read(b, binary.BigEndian, &m.Heartbeat)
	return m.Address.FromBytes(b)
}

func (m *MemberAddress) FromBytes(b *bytes.Buffer) error {
	binary.Read(b, binary.BigEndian, &m.IP)
	binary.Read(b, binary.BigEndian, &m.UDPPort)
	err := binary.Read(b, binary.BigEndian, &m.TCPPort)
	return err
}
