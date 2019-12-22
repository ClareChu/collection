package db

import (
	"io/ioutil"
	"log"
)

func ReadConfig(fileName string) ([]byte, error) {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	return yamlFile, err
}
