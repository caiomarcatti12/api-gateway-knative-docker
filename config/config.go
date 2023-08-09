package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadConfig() error {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return err
	}

	if len(config.Routes) == 0 {
		return errors.New("No routes found in the config file")
	}

	// Populando o RouteStore com as rotas carregadas
	for _, route := range config.Routes {
		GetRouteStore().Add(route)
	}

	return nil
}
