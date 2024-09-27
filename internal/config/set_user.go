package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

func (c *Config) SetUser(username string) (err error) {
	c.CurrentUserName = username

	marshalled, err := json.Marshal(c)
	if err != nil {
		return
	}

	userhome, err := os.UserHomeDir()
	if err != nil {
		return
	}

	filename := fmt.Sprintf("%v/.gatorconfig.json", userhome)
	err = os.WriteFile(filename, marshalled, fs.FileMode(os.O_WRONLY))

	return
}
