package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// LoadConfig 加载配置文件文件
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = json.Unmarshal(content, c)
	if err != nil {
		return nil, err
	}
	return c, nil
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
