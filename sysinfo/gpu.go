package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetGPU() string {
	gpu, err := readFromLspci()
	if err == nil && gpu != "" {
		return gpu
	}

	return "unknown"
}

func readFromLspci() (string, error) {
	cmd := exec.Command("sh", "-c", "lspci | grep -E 'VGA|3D'")
	cmd.Env = append(os.Environ(), "LC_ALL=C")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		lower := strings.ToLower(line)
		if strings.Contains(lower, "advanced micro devices") || strings.Contains(lower, "amd/ati") || strings.Contains(lower, "ati ") {
			return "amdgpu", nil
		}
		if strings.Contains(lower, "nvidia") {
			return "nvidia", nil
		}
		if strings.Contains(lower, "intel") {
			return "intel", nil
		}
	}

	return "", fmt.Errorf("no recognized GPU found")
}
