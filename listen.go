package main

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

func listenClipboardExceptWayland(outChan chan string) {
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for {
		b, ok := <-ch
		if !ok {
			log.Fatal("clipboard closed the channel")
		}
		outChan <- strings.TrimSpace(string(b))
	}
}

func listenClipboardWayland(outChan chan string) {
	var prev []byte
	okStderr := []byte("Clipboard content is not available as requested type \"text/plain\"")
	for {
		cmd := exec.Command("wl-paste", "-t", "text/plain")
		o, err := cmd.CombinedOutput()
		if err != nil {
			if bytes.Contains(o, okStderr) {
				continue
			} else {
				log.Fatal(err)
			}
		}
		if !bytes.Equal(prev, o) {
			prev = o
			outChan <- strings.TrimSpace(string(o))
		}
		time.Sleep(time.Second / 2)
	}
}

func ListenClipboard(outChan chan string) error {
	goos := runtime.GOOS
	log.Printf("Current OS: %s\n", goos)
	if goos == "linux" && os.Getenv("XDG_SESSION_TYPE") == "wayland" {
		log.Println("Wayland detected, use wl-paste to read clipboard instead")
		if _, err := exec.LookPath("wl-paste"); err != nil {
			return errors.New("wl-paste is not in the PATH")
		} else {
			go listenClipboardWayland(outChan)
		}
	} else {
		go listenClipboardExceptWayland(outChan)
	}
	return nil
}
