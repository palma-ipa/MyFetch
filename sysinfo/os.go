package sysinfo

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

type UserInfo struct {
	Username string
	Hostname string
}

func GetUserInfo() (UserInfo, error) {
	u, err := user.Current()
	if err != nil {
		return UserInfo{}, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return UserInfo{
		Username: u.Username,
		Hostname: hostname,
	}, nil
}

func GetOSName() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		if runtime.GOOS == "linux" {
			return "Linux"
		}
		return runtime.GOOS
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			val := strings.TrimPrefix(line, "PRETTY_NAME=")
			val = strings.Trim(val, "\"")
			val = strings.TrimSpace(val)
			if val != "" {
				return val
			}
		}
	}

	return runtime.GOOS
}

func GetKernelVersion() string {
	data, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return runtime.GOOS
	}

	return strings.TrimSpace(string(data))
}

func GetUptime() string {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "unknown"
	}

	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		return "unknown"
	}

	var seconds float64
	_, err = fmt.Sscanf(fields[0], "%f", &seconds)
	if err != nil {
		return "unknown"
	}

	hours := int(seconds / 3600)
	minutes := int((seconds - float64(hours)*3600) / 60)

	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}

func GetShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}
	return shellBaseName(shell)
}

func shellBaseName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
