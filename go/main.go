package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/load"
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

func cpu() string {
	ms, err := load.Avg()
	if err != nil {
		return "err"
	}

	return fmt.Sprintf("%.2f", ms.Load1)
}

func ram() string {
	mu, err := mem.VirtualMemory()
	if err != nil {
		return "err"
	}

	return fmt.Sprintf("%.f", mu.UsedPercent)
}

func main() {
	fmt.Printf(
		"[%scpu|%s%%ram] %s",
		cpu(),
		ram(),
		date(),
	)
}
