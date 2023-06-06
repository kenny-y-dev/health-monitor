package monitor

import (
	"testing"

	"github.com/go-ping/ping"
)

func TestCheckSuccess(t *testing.T) {
	testStats := ping.Statistics{
		PacketsRecv: 2,
		PacketsSent: 3,
	}
	if !CheckSuccessPing(testStats, false) {
		t.Errorf("Expected CheckSuccess to return True, returned False")
	}
	if CheckSuccessPing(testStats, true) {
		t.Errorf("Expected CheckSuccess to return False, returned True")
	}
	*&testStats.PacketsRecv = 0
	if CheckSuccessPing(testStats, false) {
		t.Errorf("Expected CheckSuccess to return False, returned True")
	}

}
