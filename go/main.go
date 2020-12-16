package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/host"
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

func temp() string {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		return "err"
	}

	// calculate average temperature of active sensors.
	activeSensors := 0.0
	temp := 0.0
	for _, sensor := range sensors {
		if sensor.Temperature != 0 {
			temp += sensor.Temperature
			activeSensors++
		}
	}

	return fmt.Sprintf("%.1f\u00B0C", (temp/activeSensors))
}

func main() {
	fmt.Printf(
		"%s [%scpu|%s%%ram|%s]\n",
		date(),
		cpu(),
		ram(),
		temp(),
	)
}
