package main

import (
	"fmt"
	"io/ioutil"
	"unicode"

	"gopkg.in/yaml.v2"
)

type segment struct {
	ID         int         `yaml:"id"`
	Name       string      `yaml:"name"`
	Conditions []condition `yaml:"conditions"`
}

type condition map[string][]operation

type operation map[string][]value

type value string

func (s *segment) GetConditions() []condition {
	return s.Conditions
}

func (s *segment) GetName() string {
	return s.Name
}

func (s *segment) GetID() int {
	return s.ID
}

func (v *value) IsNum() bool {
	for _, c := range *v {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func main() {
	var segments []segment
	file, _ := ioutil.ReadFile("segments.yaml")
	if err := yaml.Unmarshal(file, &segments); err != nil {
		panic(err)
	}
	fmt.Println(segments[0].GetName())
}
