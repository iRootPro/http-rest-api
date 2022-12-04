package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/iRootPro/http-rest-api/internal/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	config := apiserver.NewConfig()
	flag.Parse()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.NewAPIServer(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
