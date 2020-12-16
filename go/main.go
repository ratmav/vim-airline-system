package main

import (
    "fmt"

    "github.com/shirou/gopsutil/v3/mem"
)

func main() {
    v, _ := mem.VirtualMemory()

    fmt.Printf("RAM: %d%%\n", int(v.UsedPercent))
}
