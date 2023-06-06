package monitor

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func SendPing(target string, timeout string) (*ping.Pinger, error) {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		return pinger, fmt.Errorf("Pinger creation failed with error: %w", err)
	}

	pinger.Count = 3
	pinger.Timeout, err = time.ParseDuration(timeout)
	if err != nil {
		return pinger, fmt.Errorf("Timeout parsing failed with error: %w", err)
	}
	err = pinger.Run()
	if err != nil {
		return pinger, fmt.Errorf("Unable to run ping: %w", err)
	}
	return pinger, nil
}

func CheckSuccessPing(stats ping.Statistics, strict bool) bool {
	if stats.PacketsRecv == 0 {
		return false
	}
	if strict && stats.PacketsSent != stats.PacketsRecv {
		return false
	} else {
		return true
	}

}
