package main

import (
	"log"

	y "github.com/chrisbradleydev/go-yaml/pkg/yaml"
	"gopkg.in/yaml.v3"
)

type Values struct {
	Global struct {
		Name    string `yaml:"name"`
		Age     int    `yaml:"age"`
		Address struct {
			City  string `yaml:"city"`
			State string `yaml:"state"`
		} `yaml:"address"`
	} `yaml:"global"`
}

func main() {
	var v Values
	yamlData := []byte(`
global:
  name: Terrance Yeakey
  age: 30
  address:
    city: Oklahoma City
    state: OK
  extra1: 1
  extra2: 2
  extra3: 3`)
	err := yaml.Unmarshal(yamlData, &v)
	if err != nil {
		log.Fatal(err)
	}

	filename := "data/values.yaml"
	paths := []string{
		"global.age",
		"global.address.city",
		"global.extra2",
		"global.extra3",
	}
	y.FindAndDelete(filename, paths, yamlData)
}
