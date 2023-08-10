# API Gateway Knative-Docker

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## Introdução

O API Gateway Knative-Docker é uma solução inovadora que simula a funcionalidade do Knative no Docker. Inspirado na capacidade do Knative de gerenciar aplicações serverless no Kubernetes, este projeto visa fornecer uma alternativa para aqueles que utilizam o Docker em diferentes ambientes, seja produção, homologação ou desenvolvimento.

## Características Principais

- **Gerenciamento Dinâmico de Contêineres**: O API Gateway é capaz de iniciar contêineres dinamicamente com base nas solicitações recebidas. Se um contêiner estiver offline, ele será iniciado automaticamente.

- **Eficiência de Recursos**: Contêineres que não recebem solicitações por um período de tempo configurável são desligados, otimizando o uso de recursos.

- **Roteamento Inteligente**: O API Gateway gerencia o roteamento de solicitações para o contêiner apropriado, garantindo uma resposta rápida e eficiente.

## Configuração e Uso

1. **Configuração de Rotas**: Utilize o arquivo `config.example.yaml` para definir suas rotas, especificando detalhes como caminho, serviço associado, TTL e outros parâmetros relevantes.

2. **Execução**: Inicie o API Gateway com o comando `go run main.go`. Por padrão, o servidor será iniciado na porta 8080.

3. **Solicitações**: Envie suas solicitações HTTP para o API Gateway. Ele cuidará do roteamento e do gerenciamento dos contêineres Docker para você.

## Licença

Este projeto está licenciado sob a licença Apache 2.0. Consulte o arquivo [LICENSE](LICENSE) para obter detalhes.
