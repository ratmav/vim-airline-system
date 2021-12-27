package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func date() string {
	l, err := time.LoadLocation("UTC")
	if err != nil {
		return "err"
	}

	t := time.Now().In(l)
	return fmt.Sprint(strings.ToLower(t.Format("Jan 02 15:04 MST")))
}

func cpu_usage() string {
	// get utilization over the last 5 seconds.
	cu, err := cpu.Percent((5 * time.Second), false)
	if err != nil {
		return "err"
	}

	return fmt.Sprintf("%.f", cu[0])
}

func ram_usage() string {
	mu, err := mem.VirtualMemory()
	if err != nil {
		return "err"
	}

	return fmt.Sprintf("%.f", mu.UsedPercent)
}

func main() {
	fmt.Printf(
		"[%s%%cpu|%s%%ram] %s",
		cpu_usage(),
		ram_usage(),
		date(),
	)
}
