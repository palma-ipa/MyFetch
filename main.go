package main

import (
	"fmt"

	"myfetch/sysinfo"
)

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[1;36m"
	colorGreen  = "\033[1;32m"
)

func main() {
	fmt.Println(colorCyan + `█   █ █   █ █████ █████ █████  ███  █   █ 
██ ██  █ █  █     █       █   █     █   █ 
█ █ █   █   ████  ████    █   █     █████ 
█   █   █   █     █       █   █     █   █ 
█   █   █   █     █████   █    ███  █   █ ` + colorReset)
	fmt.Println()

	user, _ := sysinfo.GetUserInfo()
	fmt.Printf("%sUser%s: %s@%s\n", colorCyan, colorReset, user.Username, user.Hostname)

	osName := sysinfo.GetOSName()
	fmt.Printf("%sOS%s: %s\n", colorCyan, colorReset, osName)

	kernel := sysinfo.GetKernelVersion()
	fmt.Printf("%sKernel%s: %s\n", colorCyan, colorReset, kernel)

	uptime := sysinfo.GetUptime()
	fmt.Printf("%sUptime%s: %s\n", colorCyan, colorReset, uptime)

	shell := sysinfo.GetShell()
	fmt.Printf("%sShell%s: %s\n", colorCyan, colorReset, shell)

	cpu := sysinfo.GetCPU()
	fmt.Printf("%sCPU%s: %s\n", colorCyan, colorReset, cpu)

	gpu := sysinfo.GetGPU()
	fmt.Printf("%sGPU%s: %s\n", colorGreen, colorReset, gpu)

	mem := sysinfo.GetMemory()
	fmt.Printf("%sMemory%s: %s\n", colorCyan, colorReset, mem)
}
