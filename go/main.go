package main

import (
    "fmt"
    "strings"
    "time"

    //"github.com/shirou/gopsutil/v3/host"
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
	return "TODO"
}

func ram() string {
	mu, err := mem.VirtualMemory()
	if err != nil {
		return "err"
	}

	return fmt.Sprint(int(mu.UsedPercent))
}

func temp() string {
	/*
	dc, err : host.TemperatureStat()
	if err != nil {
		return "ERR"
	}

	return fmt.Sprintf("%s\u00B0C", dc)
	*/
	return "wip"
}

func main() {
	fmt.Printf(
		"%s [cpu%s%%|ram%s%%|%s]\n",
		date(),
		cpu(),
		ram(),
		temp(),
	)
}
