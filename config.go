package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title       string   `yaml:"title"`
	SubTitle    string   `yaml:"subtitle"`
	Website     string   `yaml:"website"`
	Description string   `yaml:"description"`
	Author      string   `yaml:"author"`
	Email       string   `yaml:"email"`
	Feed        string   `yaml:"feed"`
	Image       string   `yaml:"image"`
	Explicit    bool     `yaml:"explicit"`
	Categories  []string `yaml:"categories"`
	Episodes    []struct {
		Title       string `yaml:"title"`
		Link        string `yaml:"link"`
		Description string `yaml:"description"`
		Image       string `yaml:"image"`
		Size        int64  `yaml:"size"`
		Duration    int64  `yaml:"duration"`
	} `yaml:"episodes"`
}

func (c *Config) ReadFromFile(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read file at path \"%s\"", path)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func NewConfigFromFile(path string) *Config {
	var c Config
	return c.ReadFromFile(path)
}
