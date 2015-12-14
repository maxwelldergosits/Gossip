package gossip

import "gossip/members"
import "time"

type GossipMessage struct {
	To      members.GossipMember
	From    members.GossipMember
	Payload []byte
}

func RunGossip(cxt GossipContext) {

	for {

		select {
		/* normal state stuff*/
		case _ = <-time.Tick(cxt.Conf().RoundLength):
			cxt.MemberHandler().DeleteExpired(cxt.Conf().Self.Heartbeat, cxt.Conf().ExpireThreshold)
			cxt.MemberHandler().MarkSuspected(cxt.Conf().Self.Heartbeat, cxt.Conf().SuspectedThreshold)
			cxt.Conf().Self.Heartbeat += 1
			SendRoundMessage(cxt)

			/* state sync */
		case _ = <-time.Tick(cxt.Conf().SyncLength):
			RequestSync(cxt)

		/* handle new message as they come in */
		case gossip := <-cxt.Inbound():
			if gossip.Type == SyncRequest || gossip.Type == JoinRequest {
				cxt.Outbound() <- Sync(cxt, gossip.Message.From, DataMessage)
			}

			if gossip.Message.To == cxt.Conf().Self {
				cxt.ReceivedMessages() <- gossip.Message
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
			}

		}
	}
}
