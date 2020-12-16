package main

import (
    "fmt"

    //"github.com/shirou/gopsutil/v3/host"
    "github.com/shirou/gopsutil/v3/mem"
)

func date() string {
	return "TODO"
}

func cpu() string {
	return "TODO"
}

func ram() string {
	mu, err := mem.VirtualMemory()
	if err != nil {
		return "ERR"
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
	return "WIP"
}

func main() {
	fmt.Printf(
		"%s CPU %s%% RAM %s%% %s\n",
		date(),
		cpu(),
		ram(),
		temp(),
	)
}
