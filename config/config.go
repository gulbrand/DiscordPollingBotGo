package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token        string
	BotPrefix    string
	MinCarryOver int
	Emotes       string

	config *configStruct
)

type configStruct struct {
	Token        string `json:"Token"`
	BotPrefix    string `json:"BotPrefix"`
	MinCarryOver int    `json:"MinCarryOver"`
	Emotes       string `json:"Emotes"`
}

func ReadConfig() error {
	fmt.Println("Read config")

	file, err := ioutil.ReadFile("./config/config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	MinCarryOver = config.MinCarryOver
	Emotes = config.Emotes

	return nil
}
