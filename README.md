# myfetch

A lightweight, blazingly fast system fetch tool for Linux written in Go, featuring a clean ASCII banner and no bloated heavy graphics.

## Features

- **User** – current username and hostname
- **OS** – Linux distribution name from `/etc/os-release`
- **Kernel** – kernel release version
- **Uptime** – formatted uptime in hours and minutes
- **Shell** – user’s default shell
- **CPU** – processor model name
- **GPU** – detected GPU vendor (AMD/NVIDIA/Intel)
- **Memory** – used / total RAM with percentage

## Installation

```bash
git clone https://github.com/<your-username>/myfetch.git
cd myfetch
go build -ldflags="-s -w" -o myfetch main.go
sudo cp myfetch /usr/local/bin/
```

## Usage

```bash
myfetch
```
