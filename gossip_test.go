package gossip

import "testing"
import "time"

type TestContext int

func (t *TestContext) RoundLength() time.Duration { return 1 }
func (t *TestContext) SyncLength() time.Duration  { return 1 }
func (t *TestContext) RoundSize() uint            { return 10 }

func Test(t *testing.T) {

}
