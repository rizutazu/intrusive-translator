package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ollamaTranslate(input string) (string, error) {

	req := map[string]any{
		"model":  "zongwei/gemma3-translator:4b",
		"prompt": fmt.Sprintf("Translate from English to Chinese: %s", input),
		"stream": false,
	}
	b, _ := json.Marshal(req)
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	var d map[string]any
	{
		if resp.StatusCode != 200 {
			return "", fmt.Errorf("ollama status code %d", resp.StatusCode)
		}
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		err = json.Unmarshal(b, &d)
		if err != nil {
			return "", err
		}
	}

	output, ok := d["response"]
	if !ok {
		return "", errors.New("response does not contain response field")
	}
	s, ok := output.(string)
	if !ok {
		return "", errors.New("response field is not string")
	}
	return strings.TrimSpace(s), nil
}
