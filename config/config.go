package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadConfig() WxConfig {
	data, err := ioutil.ReadFile("./config/wx_config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	var wxConfig WxConfig
	err = yaml.Unmarshal(data, &wxConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wxConfig)
	return wxConfig
}

type WxConfig struct {
	CorpId      string `yaml:"CorpId"`
	AgentId     int    `yaml:"AgentId"`
	AgentSecret string `yaml:"AgentSecret"`
	CronExp     string `yaml:"CronExp"`
}
