package runner

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var currentCmd *exec.Cmd

func Run(file string) {
	startProcess(file)
}

func Restart(file string) {
	if currentCmd != nil && currentCmd.Process != nil {
		fmt.Println("🔁 Restarting due to file change...")
		currentCmd.Process.Kill()
	}
	startProcess(file)
}

func startProcess(file string) {
	cmd := exec.Command("go", "run", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	currentCmd = cmd

	go func() {
		if err := cmd.Run(); err != nil {
			fmt.Println("❌ Process exited with error:", err)
		}
	}()
}

func init() {
	// Handle Ctrl+C to clean up
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if currentCmd != nil && currentCmd.Process != nil {
			currentCmd.Process.Kill()
		}
		os.Exit(0)
	}()
}
