package gossip

import "gossip/members"
import "time"

type GossipMessage struct {
	To      members.GossipMember
	From    members.GossipMember
	Payload []byte
}

func StartGossip(conf GossipConf) (received, outbound chan GossipMessage, err error) {

	// open files, listen on sockets, etc
	cxt, err := createContext(conf, received, outbound)
	if err != nil {
		return received, outbound, err
	}

	go startGossip(cxt)

	return received, outbound, nil
}

func startGossip(cxt GossipContext) {

	for {

		select {
		/* normal state stuff*/
		case _ = <-time.Tick(cxt.Conf().RoundLength):
			cxt.Outbound() <- SendRoundMessage(cxt)

			/* state sync */
		case _ = <-time.Tick(cxt.Conf().SyncLength):
			cxt.Outbound() <- RequestSync(cxt)

		/* handle new message as they come in */
		case gossip := <-cxt.Inbound():
			if gossip.Type == SyncRequest {
				cxt.Outbound() <- Sync(cxt, gossip.Message.From, DataMessage)
			}

			if gossip.Type == JoinRequest {
				cxt.Outbound() <- Sync(cxt, gossip.Message.From, DataMessage)
			}

			if gossip.Message.To == cxt.Conf().Self {
				cxt.ReceivedMessages() <- gossip.Message
			} else {
				//forwardMessage(cxt, message)
			}

			members.UpdateMember(cxt.MemberHandler(), gossip.Message.From, cxt.Round())
			for _, member := range gossip.Members {
				members.UpdateMember(cxt.MemberHandler(), member, cxt.Round())
			}

		/* send a message from the localhost */
		case msg := <-cxt.OutboundMessages():
			cxt.Outbound() <- Gossip{
				Type:    DataMessage,
				Message: msg,
				Members: []members.GossipMember{},
			}

		}
	}
}
