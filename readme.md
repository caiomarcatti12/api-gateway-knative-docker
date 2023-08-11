# API Gateway Knative-Docker

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE) ![Static Badge](https://img.shields.io/badge/N%C3%A3o%20pronto%20para%20produ%C3%A7%C3%A3o-red)

## Introdução

O API Gateway Knative-Docker é uma solução inovadora que simula a funcionalidade do Knative no Docker. Inspirado na capacidade do Knative de gerenciar aplicações serverless no Kubernetes, este projeto visa fornecer uma alternativa para aqueles que utilizam o Docker em diferentes ambientes, seja produção, homologação ou desenvolvimento.

## Características Principais

- **Gerenciamento Dinâmico de Contêineres**: O API Gateway é capaz de iniciar contêineres dinamicamente com base nas solicitações recebidas. Se um contêiner estiver offline, ele será iniciado automaticamente.

- **Eficiência de Recursos**: Contêineres que não recebem solicitações por um período de tempo configurável são desligados, otimizando o uso de recursos.

- **Roteamento Inteligente**: O API Gateway gerencia o roteamento de solicitações para o contêiner apropriado, garantindo uma resposta rápida e eficiente.

## Configuração das Rotas do API Gateway

O API Gateway do nosso projeto utiliza uma estrutura específica para configurar e gerenciar rotas que redirecionam solicitações para contêineres Docker específicos. Essa configuração aborda aspectos como o caminho da rota, o serviço de destino, tentativas de retentativa, checagem de saúde e mais.

Para entender completamente como configurar e o comportamento esperado dessas rotas, consulte o guia detalhado disponível em [Configuração de Rotas](./ROUTE_CONFIGURATION.md).

## Inicialização do Ambiente de Desenvolvimento

Para configurar e iniciar o ambiente de desenvolvimento do projeto, consulte o guia [Desenvolvimento](./desenvolvimento.md).

## Inicialização do Ambiente de Produção

Para configurar e iniciar o ambiente de produção do projeto, consulte o guia [Produção](./producao.md).

## Como Contribuir

Estamos sempre abertos a contribuições! Se você deseja ajudar a melhorar o projeto, seja através de correções de bugs, melhorias ou novas funcionalidades, siga nosso [Guia de Contribuição](CONTRIBUTING.md) para entender o processo e garantir que sua contribuição seja integrada da melhor forma possível.

## Código de Conduta

Estamos comprometidos em proporcionar uma comunidade acolhedora e inclusiva para todos. Esperamos que todos os participantes do projeto sigam nosso [Código de Conduta](CODE_OF_CONDUCT.md). Pedimos que leia e siga estas diretrizes para garantir um ambiente respeitoso e produtivo para todos os colaboradores.

## Licença

Este projeto está licenciado sob a licença Apache 2.0. Consulte o arquivo [LICENSE](LICENSE) para obter detalhes.
