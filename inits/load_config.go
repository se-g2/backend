package inits

import (
	"backend/global"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfig(filename string) error {

	data, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &global.Config)

}
