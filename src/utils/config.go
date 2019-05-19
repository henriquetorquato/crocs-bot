package utils

import (
	"encoding/json"
	"os"
)

const configFile = "./config.json"

func read() Configuration {
	var file, err = os.Open(configFile)
	defer file.Close()

	HandleError(err)

	decoder := json.NewDecoder(file)
	config := Configuration{}
	decoder.Decode(&config)

	return config
}

// Configuration file structure definition
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
	Message struct {
		Use          string
		UseWithSocks string
		DontUse      string
	}
	Variations struct {
		Heat []string
		Cold []string
		Rain []string
	}
}

var Config = read()
