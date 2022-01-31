package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil" // used to read our json config file
)

type configStruct struct {
	BotPrefix string `json:"BotPrefix"`
	Currency  string `json:"Currency"`
}

var (
	BotPrefix string // store our bot prefix from json
	Currency  string
	config    *configStruct
)

func ReadConfig() error {
	fmt.Println("Reading config...")

	// read config.json, value is stored in file var and if an error ocurrs it will be stored in err
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Copying the value of file into var config, and if there any error we are storing it in err
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Config was copied succesfully!")

	// Store the config variables in separate global variables
	Currency = config.Currency
	BotPrefix = config.BotPrefix

	//If there isn't any error we will return nil.
	return nil
}
