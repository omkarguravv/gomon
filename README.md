# gomon 🧠⚡️

A `nodemon`-like utility for Go — auto-restarts your `go run` process when `.go` files change.

> 🔁 Save code → 🚀 Auto-restart → 🧪 Rapid Dev Flow

---

## ✨ Features

- 📂 Watches `.go` files for changes in the current directory (by default)
- 🔁 Automatically restarts your `go run` process
- ⏱️ Shows time taken to rebuild and restart
- 🔍 Optional `--exclude` to ignore directories (like `vendor`, `testdata`)
- 📜 Optional `gomon.yaml` for configuration
- 🧠 Auto-detects `go.mod` and runs `go build && ./binary` if available
- 🧼 Graceful exit handling (`Ctrl+C` stops child process too)

---

## 🚀 Installation

```bash
go install github.com/omkarguravv/gomon@latest
