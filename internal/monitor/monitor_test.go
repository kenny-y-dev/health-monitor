package monitor

import (
	"testing"

	"github.com/go-ping/ping"
)

func TestCheckPingSuccess(t *testing.T) {
	want := 3

	pinger, err := SendPing("127.0.0.1", "5s")
	if err != nil {
		t.Errorf("Error sending ping: %v", err)
	}
	if want != pinger.Statistics().PacketsRecv {
		t.Errorf("Ping module failed to receive 3 packets")
	}
}

func TestCheckPingFailure(t *testing.T) {
	want := 0
	pinger, err := SendPing("203.0.113.1", "1s") // IP address should not route per RFC 5737
	if err != nil {
		t.Errorf("Ping module failed with error: %v", err)
	}
	if want != pinger.Statistics().PacketsRecv {
		t.Errorf("Ping module received reply packets when it shouldn't have")
	}

}

func TestCheckPingError(t *testing.T) {
	pinger, err := SendPing("127.0.0.1", "5invalid")
	if pinger != nil && err == nil {
		t.Errorf("Ping module did not throw error on invalid timeout value")
	}

}
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
	testStats.PacketsRecv = 0
	if CheckSuccessPing(testStats, false) {
		t.Errorf("Expected CheckSuccess to return False, returned True")
	}

}
