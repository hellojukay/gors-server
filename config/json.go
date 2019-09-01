package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var C *Config = &Config{}

// LoadConfig 加载配置文件文件
func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, C)
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Data struct {
		Dir string `json:"dir"`
	} `json:"data"`
	LDAP struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"ldap"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}
