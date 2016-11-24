package tools

import (
	"encoding/json"
	"log"
	"os"
)

// Config holds the global configuration of moolinet.
type Config struct {
	ChallengesPath  string
	DatabasePath    string
	StaticPath      string
	CookieSecretKey string
}

// GeneralConfig for the programm
var GeneralConfig Config

// LoadGeneralConfigFromFile loads a config that will be available to the whole program
func LoadGeneralConfigFromFile(file string) error {
	reader, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func() {
		if errDefer := reader.Close(); errDefer != nil {
			log.Println("Unable to close a reader: ", errDefer.Error())
		}
	}()

	decoder := json.NewDecoder(reader)

	GeneralConfig = Config{}
	err = decoder.Decode(&GeneralConfig)

	return err
}
