package main

import "time"
import "log"
import "gossip"
import "gossip/members"
import "sync"

func main() {
	N := 10
	confs := make([]gossip.GossipConf, N, N)
	cxts := make([]gossip.GossipContext, N, N)
	inbounds := make([]chan gossip.GossipMessage, N, N)
	outbounds := make([]chan gossip.GossipMessage, N, N)
	inboundGs := make([]chan gossip.Gossip, N, N)
	outboundGs := make([]chan gossip.Gossip, N, N)
	ids := make([]members.MemberID, N, N)
	idToIndex := make(map[members.MemberID]int, 2000)
	for i := 0; i < N; i++ {
		outbounds[i] = make(chan gossip.GossipMessage, 100)
		inbounds[i] = make(chan gossip.GossipMessage, 100)
		ids[i] = members.NewID(uint64(i+1), 0)
		idToIndex[ids[i]] = i
		confs[i] = gossip.GossipConf{
			RoundLength:        time.Duration(500 * time.Millisecond),
			SyncLength:         time.Duration(500 * time.Millisecond),
			SuspectedThreshold: 1000,
			ExpireThreshold:    10000,
			RoundSize:          20,
			Self:               members.GossipMember{ID: ids[i]},
		}
		cxt, _ := gossip.CreateContext(confs[i], inbounds[i], outbounds[i])
		inboundGs[i] = cxt.InboundChannel
		outboundGs[i] = cxt.OutboundChannel
		cxts[i] = cxt
		h := members.CreateMemoryMemberHandler()
		cxt.SetMemberHandler(&h)

		go gossip.RunGossip(cxt)
	}

	for i := 0; i < N; i++ {
		inboundGs[i] <- gossip.Gossip{
			Type: gossip.JoinRequest,
			Message: gossip.GossipMessage{
				To:   members.GossipMember{ID: ids[i]},
				From: members.GossipMember{ID: ids[((i+N)-1)%N]},
			},
		}
	}
	messages := make(chan gossip.Gossip, 1000)
	go sendMessages(N, idToIndex, outboundGs, inboundGs, messages)

  var lock sync.Mutex
	members := make(map[string]members.GossipMember, len(idToIndex))

  go serve(&lock, &members)

	for msg := range messages {
    members[msg.Message.From.ID.ToString()] = msg.Message.From
		var typeString string
		if msg.Type == gossip.JoinRequest {
			typeString = "JoinRequest"
		} else if msg.Type == gossip.SyncRequest {
			typeString = "SyncRequest"
		} else {
			typeString = "DataMessage"
		}

		log.Println("message from:", msg.Message.From.ID.Upper,
			msg.Message.From.ID.Lower, "type:", typeString)
	}
}

func printStats(start time.Time, n, m, N uint64) time.Time {
	now := time.Now()
	if now.Sub(start) > 5*time.Second {
		diff := now.Sub(start)
		rate := float64(n) * float64(time.Second) / float64(diff)
		ratem := float64(m) * float64(time.Second) / float64(diff)
		log.Println("for N=", N, "msgs/sec =", rate, "members/sec =", ratem)
		n = 0
		m = 0
		return time.Now()
	}
	return start
}

func sendMessages(N int, idToIndex map[members.MemberID]int, outboundGs, inboundGs []chan gossip.Gossip, messages chan gossip.Gossip) {

	start := time.Now()
	doPrint := false
	var n uint64 = 0
	var m uint64 = 0
	for {
		for i := 0; i < N; i++ {
			select {
			case g := <-outboundGs[i]:
				messages <- g
				n += 1
				m += uint64(len(g.Members))
				inboundGs[idToIndex[g.Message.To.ID]] <- g
			default:
				if doPrint {
					start = printStats(start, n, m, uint64(N))
				}
			}
		}
	}
}
