package main

import (
	"fmt"

	"github.com/kenny-y-dev/health-monitor/internal/monitor"
	"github.com/kenny-y-dev/health-monitor/internal/notify"
)

func main() {
	fmt.Println(monitor.Send())
	notify.Send()
}
