package gossip

import "bytes"
import "encoding/binary"
import "gossip/members"

func (gm *GossipMessage) ToBytes(buf *bytes.Buffer) {
	gm.From.ToBytes(buf)
	gm.To.ToBytes(buf)
	binary.Write(buf, binary.BigEndian, uint64(len(gm.Payload)))
	buf.Write(gm.Payload)
}

func (g *Gossip) ToBytes(buf *bytes.Buffer) {
	binary.Write(buf, binary.BigEndian, uint64(g.Type))
	g.Message.ToBytes(buf)
	binary.Write(buf, binary.BigEndian, uint64(len(g.Members)))
	for _, m := range g.Members {
		m.ToBytes(buf)
	}
}

func (gm *GossipMessage) FromBytes(buf *bytes.Buffer) {
	gm.From.FromBytes(buf)
	gm.To.FromBytes(buf)
	var payloadLength uint64
	binary.Read(buf, binary.BigEndian, &payloadLength)
	gm.Payload = make([]byte, payloadLength, payloadLength)
	buf.Read(gm.Payload)
}

func (g *Gossip) FromBytes(buf *bytes.Buffer) {
	var Type uint64
	binary.Write(buf, binary.BigEndian, &Type)
	g.Type = MessageType(Type)
	g.Message.FromBytes(buf)
	var nMembers uint64
	binary.Read(buf, binary.BigEndian, uint64(len(g.Members)))
	g.Members = make([]members.GossipMember, nMembers, nMembers)
	for _, m := range g.Members {
		m.FromBytes(buf)
	}
}
