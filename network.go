package gossip

import (
	"bytes"
	"gossip/members"
	"log"
	"net"
)

const UDPMessageSize = 1492

func ListenUDP(port uint16, messages chan Gossip) error {
	addr := net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: int(port),
	}
	ln, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		return err
	}
	go func() {
		for {
			var buf = make([]byte, UDPMessageSize, UDPMessageSize)
			n, _, err := ln.ReadFromUDP(buf)
			if err != nil {
				log.Println("error reading packet:", err.Error())
			}
			go func() {
				var buffer bytes.Buffer
				buffer.Write(buf[0:n])
				var g Gossip
				g.FromBytes(&buffer)
				messages <- g
			}()
		}
	}()
	return nil
}

func SendUDP(local members.MemberAddress, messages chan Gossip) error {

	localAddr := net.UDPAddr{
	  IP:   net.IPv4(local.IP[0], local.IP[1], local.IP[2], local.IP[3]),
		Port: 0,
	}

	go func() {
		for g := range messages {
			remote := g.Message.To.Address

			remoteAddr := net.UDPAddr{
				IP:   net.IPv4(remote.IP[0], remote.IP[1], remote.IP[2], remote.IP[3]),
				Port: int(remote.UDPPort),
			}

			conn, err := net.DialUDP("udp", &localAddr, &remoteAddr)
			if err != nil {
				log.Println("error reading packet:", err.Error())
			}
			var buffer bytes.Buffer
			g.ToBytes(&buffer)
			conn.Write(buffer.Bytes())
			defer conn.Close()
		}
	}()
	return nil
}
