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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func LoadConfig() error {
	// Em produção, esperamos que o arquivo config.yaml esteja na raiz
	executable, err := os.Executable()
	if err != nil {
		log.Fatal("Erro ao obter o caminho do binário:", err)
	}
	execDir := filepath.Dir(executable)

	// Define o caminho absoluto para o arquivo .env na mesma pasta do binário
	envPath := filepath.Join(execDir, "config.yaml")

	configFile, err := ioutil.ReadFile(envPath)
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
