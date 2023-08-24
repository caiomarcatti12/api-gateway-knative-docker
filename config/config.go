/*
 * Copyright 2023 Caio Matheus Marcatti Calimério
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"api-gateway-knative-docker/cors"
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

	cors.GetCorsStore().Add(config.CORS)

	return nil
}
