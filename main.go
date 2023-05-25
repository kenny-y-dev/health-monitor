package main

import (
	"log"

	"github.com/kenny-y-dev/health-monitor/internal/config"
	"github.com/kenny-y-dev/health-monitor/internal/monitor"
	"github.com/kenny-y-dev/health-monitor/internal/notify"
)

func main() {
	status := true
	cfg := config.Build()
	pinger, err := monitor.SendPing(cfg.MonitorTarget, cfg.MonitorTimeout)
	if err != nil {
		log.Fatalf("Ping module failed %v", err)
	}
	up := monitor.CheckSuccessPing(*pinger.Statistics(), cfg.MonitorCheckStrict)
	if up != status {
		status = !status
		if !status {
			log.Printf("Target failed monitor check")
			notify.SendFailure(cfg.NotifyTarget)
			// target down
		}
		if status {
			log.Printf("Target healthy again")
			notify.SendSuccss(cfg.NotifyTarget)
		}
	}
}
