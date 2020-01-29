package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var Config map[string]string

func load() error {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}
	log.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	jsonBlob, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var config map[string]string
	err = json.Unmarshal(jsonBlob, &config)
	if err != nil {
		return err
	}

	if len(config["db_path"]) == 0 {
		config["db_path"] = "./data/all.data.tieba.baidu.com.db"
	}
	if len(config["img_dir"]) == 0 {
		config["img_dir"] = "./data/IMG"
	}
	log.Println(config)
	Config = config
	return nil
}

func GetConfig() (map[string]string, error) {
	var err error
	if nil == Config {
		err = load()
	}
	if err != nil {
		return nil, err
	}
	return Config, nil
}

func Get(key string) string {
	config, err := GetConfig()
	value := ""
	if err == nil {
		if len(config[key]) > 0 {
			value = config[key]
		}
	}
	return value
}
