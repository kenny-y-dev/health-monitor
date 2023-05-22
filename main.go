package main

import (
	"fmt"

	"github.com/kenny-y-dev/health-monitor/internal/config"
	"github.com/kenny-y-dev/health-monitor/internal/monitor"
	"github.com/kenny-y-dev/health-monitor/internal/notify"
)

func main() {
	cfg := config.Build()
	fmt.Println(monitor.Send(cfg.MonitorHost))
	notify.Send()
}
