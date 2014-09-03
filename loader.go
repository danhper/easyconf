package easyconf

import (
	"encoding/json"
	"gopkg.in/yaml.v1"
	"io"
	"io/ioutil"
)

type Parser interface {
	ParseConfig(config io.Reader) (map[string]interface{}, error)
}

type JsonParser struct{}

func (p *JsonParser) ParseConfig(config io.Reader) (result map[string]interface{}, err error) {
	err = json.NewDecoder(config).Decode(&result)
	return
}

type YamlParser struct{}

func (p *YamlParser) ParseConfig(config io.Reader) (result map[string]interface{}, err error) {
	var data []byte
	if data, err = ioutil.ReadAll(config); err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &result)
	return
}
