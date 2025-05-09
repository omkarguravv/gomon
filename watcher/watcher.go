package watcher

import (
	"log"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/omkarguravv/gomon/utils"
)

func Watch(path string, onChange func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	debounce := utils.NewDebouncer(500 * time.Millisecond)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Only reload on file write or create
				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove|fsnotify.Rename) != 0 {
					if filepath.Ext(event.Name) == ".go" {
						debounce.Call(onChange)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal("Failed to watch:", err)
	}
	<-done
}
