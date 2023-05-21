package monitor

import (
	"log"
	"time"

	"github.com/go-ping/ping"
)

func Send() ping.Statistics {
	pinger, err := ping.NewPinger("8.8.8.8")
	if err != nil {
		log.Fatalf("Pinger creation failed with error: %v", err)
	}

	pinger.Count = 3
	pinger.Timeout, err = time.ParseDuration("5s")
	if err != nil {
		log.Fatalf("Timeout parsing failed with error: %v", err)
	}
	err = pinger.Run()
	if err != nil {
		log.Fatalf("Ping failed with error: %v", err)
	}
	return *pinger.Statistics()
}

func CheckSuccess(stats ping.Statistics, strict bool) bool {
	if stats.PacketsRecv == 0 {
		return false
	}
	if strict && stats.PacketsSent != stats.PacketsRecv {
		return false
	} else {
		return true
	}

}
