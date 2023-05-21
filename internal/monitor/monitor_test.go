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
	if !CheckSuccess(testStats, false) {
		t.Errorf("Expected CheckSuccess to return True, returned False")
	}
	if CheckSuccess(testStats, true) {
		t.Errorf("Expected CheckSuccess to return False, returned True")
	}
	*&testStats.PacketsRecv = 0
	if CheckSuccess(testStats, false) {
		t.Errorf("Expected CheckSuccess to return False, returned True")
	}

}
