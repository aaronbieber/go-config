package config

import (
	"gopkg.in/gcfg.v1"
	"log"
)

type Config struct {
	Main struct {
		Key    string
		Host   string
		Port   int
		CmdDir string
	}
}

var cfg Config

func loadConfig(loadPath string) {
	err := gcfg.ReadFileInto(&cfg, loadPath)

	if err != nil {
		log.Fatal(err)
		panic("Error loading configuration.")
	}
}

func Conf() Config {
	loadConfig("config.gcfg")
	return cfg
}

func ConfFrom(loadPath string) Config {
	loadConfig(loadPath)
	return cfg
}
