package main

import (
	"log"

	"github.com/kenny-y-dev/health-monitor/internal/config"
	"github.com/kenny-y-dev/health-monitor/internal/monitor"
	"github.com/kenny-y-dev/health-monitor/internal/notify"
)

func main() {
	cfg := config.Build()
	cfg.PrintConfig()
	status := true
	for {
		checkStatus(cfg, &status)
	}
}

func checkStatus(cfg config.MonitorConfig, status *bool) {
	pinger, err := monitor.SendPing(cfg.MonitorTarget, cfg.MonitorTimeout)
	if err != nil {
		log.Fatalf("Ping module failed %v", err)
	}
	up := monitor.CheckSuccessPing(*pinger.Statistics(), cfg.MonitorCheckStrict)
	if up != *status {
		*status = !*status
		if !*status {
			log.Printf("Target failed monitor check")
			res, err := notify.HttpReq(cfg.NotifyMethod, cfg.NotifyTarget, cfg.NotifyDownJSON)
			if err != nil {
				log.Printf("Failed to notify target down with error: %v", err)
			} else {
				log.Printf("Sent Down notification to target, return code: %v", res.StatusCode)
			}
			// target down
		}
		if *status {
			log.Printf("Target healthy")
			res, err := notify.HttpReq(cfg.NotifyMethod, cfg.NotifyTarget, cfg.NotifyUpJSON)
			if err != nil {
				log.Printf("Failed to notify target down with error: %v", err)
			} else {
				log.Printf("Sent Up notification to target, return code: %v", res.StatusCode)
			}
		}
	}
}
