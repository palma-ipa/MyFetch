package sysinfo

import (
	"bufio"
	"os"
	"strings"
)

func GetCPU() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	modelName := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "model name\t:") {
			modelName = strings.TrimPrefix(line, "model name\t:")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "unknown"
	}

	if modelName == "" {
		return "unknown"
	}

	modelName = strings.Join(strings.Fields(modelName), " ")
	modelName = strings.TrimSpace(modelName)

	vendorStrings := []string{
		"Intel Corporation",
		"Advanced Micro Devices, Inc.",
		"AMD",
		"ARM",
	}
	for _, vs := range vendorStrings {
		modelName = strings.TrimPrefix(modelName, vs+" ")
	}

	modelName = strings.TrimSpace(modelName)
	modelName = strings.Trim(modelName, "\"")

	if modelName == "" {
		return "unknown"
	}
	return modelName
}
