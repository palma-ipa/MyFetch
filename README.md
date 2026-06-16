# myfetch

A lightweight and fast system fetch tool for Linux written in Go. 
No useless bloat, just clean, instant system specs.

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

## Example
<img width="2560" height="1440" alt="image" src="https://github.com/user-attachments/assets/60aea38f-a13f-41a3-960e-c9449705db01" />

  
## Installation & Setup

```bash
git clone [https://github.com/palma-ipa/myfetch.git](https://github.com/palma-ipa/myfetch.git)
cd myfetch
go build -ldflags="-s -w" -o myfetch main.go
sudo cp myfetch /usr/local/bin/

Usage

myfetch
