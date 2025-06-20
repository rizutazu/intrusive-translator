package main

import (
	"github.com/pemistahl/lingua-go"
)

var detector lingua.LanguageDetector

func init() {
	languages := []lingua.Language{
		lingua.English,
		lingua.Japanese,
		lingua.Korean,
		lingua.Chinese,
	}
	detector = lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()
}

func isEnglish(input string) bool {
	if language, _ := detector.DetectLanguageOf(input); language == lingua.English {
		return true
	}
	return false
}
