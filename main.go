package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Ltime)
}

func mainLoop(inChan chan string) {
	for {
		input, ok := <-inChan
		if !ok {
			return
		}
		if isEnglish(input) {
			log.Printf("New valid input: %s\n", input)
			translated, err := ollamaTranslate(input)
			if err != nil {
				log.Printf("Failed to translate: %s\n", err)
				continue
			}
			log.Printf("Translated => %s\n", translated)
			err = pushNotification(input, translated)
			if err != nil {
				log.Printf("Failed to push notification: %s\n", err)
			}
		}
	}
}

func main() {

	ch := make(chan string)
	err := ListenClipboard(ch)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Started")
	mainLoop(ch)
}
