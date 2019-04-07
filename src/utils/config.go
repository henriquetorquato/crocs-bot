package utils

import (
	"encoding/json"
	"os"
)

const configFile = "../config.json"

func read() Configuration {
	var file, _ = os.Open(configFile)
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}
	decoder.Decode(&config)

	return config
}

type Configuration struct {
	Advisor struct {
		Token string
	}
	Facebook struct {
		Page struct {
			ID    string
			Token string
		}
	}
	Twitter struct {
		Consumer struct {
			Key    string
			Secret string
		}
		Access struct {
			Token  string
			Secret string
		}
	}
}

var Config = read()
