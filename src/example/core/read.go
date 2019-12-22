package core

import (
	"fmt"
	"github.com/ClareChu/collection/src/example/copiey"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Config struct {
	FileName      string   `json:"fileName"`
	Include       []string `json:"include"`
	Profiles      string   `json:"profiles"`
	DefaultConfig []byte   `json:"defaultConfig"`
	ProfileConfig []byte   `json:"profileConfig"`
}

func NewConfig(fileName string) *Config {
	profile := os.Getenv("APP_PROFILES_ACTIVE")
	config := &Config{FileName: fileName,
		Profiles: profile}
	config.ReadConfig()
	return config
}

func (config *Config) ReadConfig() error {
	defaultConfig, err := ioutil.ReadFile(config.FileName)
	if err != nil {
		return err
	} else {
		config.DefaultConfig = defaultConfig
	}
	if config.Profiles == "" {
		log.Println("include  profile: default")
		return nil
	}
	log.Println("include  profile:", config.Profiles)
	profileConfig, err := ioutil.ReadFile(getProfileName(config.FileName, config.Profiles))
	if err != nil {
		return err
	} else {
		config.ProfileConfig = profileConfig
	}
	return nil
}

func (config *Config) Marshal(o interface{}) {
	toValue := yaml.Unmarshal(config.DefaultConfig, o)
	fromValue := yaml.Unmarshal(config.ProfileConfig, o)
	copiey.Copy(toValue, fromValue)

}

func getProfileName(filename, profile string) string {
	file := strings.Split(filename, ".yaml")
	return fmt.Sprintf("%v-%v.yaml", file, profile)
}
