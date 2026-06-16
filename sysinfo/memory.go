package sysinfo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetMemory() string {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return "unknown"
	}

	lines := strings.Split(string(data), "\n")

	memInfo := make(map[string]int64)
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			key := fields[0]
			if len(key) > 0 && key[len(key)-1] == ':' {
				key = key[:len(key)-1]
			}
			val, err := strconv.ParseInt(fields[1], 10, 64)
			if err == nil {
				memInfo[key] = val
			}
		}
	}

	total := memInfo["MemTotal"]
	available := memInfo["MemAvailable"]

	if total == 0 {
		return "unknown"
	}

	totalMB := float64(total) / 1024
	usedMB := float64(total-available) / 1024
	if available == 0 {
		usedMB = totalMB
	}
	percentage := (usedMB / totalMB) * 100

	if available == 0 {
		return fmt.Sprintf("%.0f MiB / %.0f MiB (%.1f%%)", usedMB, totalMB, percentage)
	}
	return fmt.Sprintf("%.0f MiB / %.0f MiB (%.1f%%)", usedMB, totalMB, percentage)
}
