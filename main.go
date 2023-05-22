package main

import (
	"github.com/kenny-y-dev/health-monitor/internal/config"
	"github.com/kenny-y-dev/health-monitor/internal/monitor"
	"github.com/kenny-y-dev/health-monitor/internal/notify"
)

func main() {
	cfg := config.Build()
	monitor.Send(cfg.MonitorHost)
	notify.Send(cfg.NotifyHost)
}
