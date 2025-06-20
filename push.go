package main

import _ "embed"
import (
	"fmt"
	"github.com/gen2brain/beeep"
)

func init() {
	beeep.AppName = "Aggressive Translator"
}

//go:embed icon.svg
var icon []byte

func pushNotification(original, translated string) error {
	format := "%s\n↑\n%s"
	message := fmt.Sprintf(format, translated, original)
	return beeep.Notify(" ", message, icon)
}
