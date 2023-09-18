package lib

import (
	"encoding/json"
	"log"
	"os"
)

func GetSecrets(secretConfig *string, config *Config) {

	configFile, errReadFile := os.ReadFile("system_configs/config.json")
	if errReadFile != nil {
		log.Fatal(errReadFile.Error())
	}

	errUnMarshall := json.Unmarshal([]byte(configFile), config)
	if errUnMarshall != nil {
		log.Fatal("Error Unmarshal config.json during startup")
	}

}
