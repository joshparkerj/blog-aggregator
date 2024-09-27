package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func Read() (config Config, err error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return
	}

	filename := fmt.Sprintf("%v/.gatorconfig.json", userHome)
	contents, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	contentsReader := bytes.NewReader(contents)
	decoder := json.NewDecoder(contentsReader)

	decoder.Decode(&config)

	return
}
