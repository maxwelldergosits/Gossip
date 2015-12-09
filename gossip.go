package gossip

import "time"

type GossipMessage struct {
  To GossipMember
  From GossipMember
  payload []byte
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

func startGossip(cxt gossipContext) {

	for {

		select {
		/* normal state stuff*/
		case _ = <-time.Tick(cxt.Conf().RoundLength):
			round(cxt)

			/* state sync */
		case _ = <-time.Tick(cxt.Conf().SyncLength):
			requestSync(cxt)

		/* handle new message as they come in */
		case msg := <-cxt.Inbound():
			handleMessage(cxt, msg)

		/* send a message from the localhost */
		case msg := <-cxt.OutboundMessages():
			sendMessage(cxt, msg)

		}
	}
}
