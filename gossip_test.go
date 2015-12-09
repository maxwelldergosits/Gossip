package gossip

import "testing"
import "time"

type TestContext int

func (t *TestContext) RoundLength() time.Duration { return 1 }

func TestTick(t *testing.T) {
	var testContext TestContext
	c := make(chan []byte)
	startGossip(&testContext, c, c)

}
