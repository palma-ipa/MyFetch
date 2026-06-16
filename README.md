# myfetch

A lightweight and fast system fetch tool for Linux written in Go. No ASCII art, no bloat—just clean, instant system specs.

---

## Dependencies

- **Go** (1.21+) - Required for compiling the source code.
- **pciutils** (`lspci`) - Required for exact GPU model detection.

---
## Displayed Specs

- User (`user@hostname`)
- OS Name
- Kernel Version
- Uptime
- Shell Name
- CPU Model
- GPU Model
- Memory Usage (`Used MiB / Total MiB (%)`)

  
## Installation & Setup

```bash
git clone [https://github.com/palma-ipa/myfetch.git](https://github.com/palma-ipa/myfetch.git)
cd myfetch
go build -ldflags="-s -w" -o myfetch main.go
sudo cp myfetch /usr/local/bin/

Usage

myfetch
