package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/omkarguravv/gomon/runner"
	"github.com/omkarguravv/gomon/watcher"
)

func Execute() {
	var fileToRun string
	var watchPath string

	flag.StringVar(&fileToRun, "f", "", "Go file to run (default: main.go in current dir)")
	flag.StringVar(&watchPath, "w", ".", "Directory to watch")
	flag.Parse()

	if fileToRun == "" && len(flag.Args()) > 0 {
		fileToRun = flag.Args()[0]
	}

	if fileToRun == "" {
		fileToRun = "main.go"
	}

	if _, err := os.Stat(fileToRun); os.IsNotExist(err) {
		fmt.Println("❌ File not found:", fileToRun)
		os.Exit(1)
	}
	absFile, _ := filepath.Abs(fileToRun)
	fmt.Println("👀 Watching", watchPath)
	fmt.Println("🚀 Running", absFile)

	go runner.Run(absFile)

	watcher.Watch(watchPath, func() {
		runner.Restart(absFile)
	})

}
