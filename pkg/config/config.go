package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Path string

func (c Path) String() string {
	return string(c)
}

// Load config
func Load(path Path, config interface{}) error {
	if len(path.String()) != 0 {
		exists, err := exist(path.String())
		if err != nil {
			return err
		}
		if !exists {
			return err
		}
	} else {

		if err := envconfig.Process("", config); err != nil {
			return err
		}
		return nil
	}

	err := godotenv.Overload(path.String())
	if err != nil {
		return err
	}
	if err := envconfig.Process("", config); err != nil {
		return err
	}
	return nil
}

func exist(path string) (exists bool, err error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}
