# API Gateway Knative-Docker

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](license) ![Static Badge](https://img.shields.io/badge/N%C3%A3o%20pronto%20para%20produ%C3%A7%C3%A3o-red)

## Introdução

O API Gateway Knative-Docker é uma solução inovadora que simula a funcionalidade do Knative no Docker. Inspirado na capacidade do Knative de gerenciar aplicações serverless no Kubernetes, este projeto visa fornecer uma alternativa para aqueles que utilizam o Docker em diferentes ambientes, seja produção, homologação ou desenvolvimento.

## Características Principais

- **Gerenciamento Dinâmico de Contêineres**: O API Gateway é capaz de iniciar contêineres dinamicamente com base nas solicitações recebidas. Se um contêiner estiver offline, ele será iniciado automaticamente.

- **Eficiência de Recursos**: Contêineres que não recebem solicitações por um período de tempo configurável são desligados, otimizando o uso de recursos.

- **Roteamento Inteligente**: O API Gateway gerencia o roteamento de solicitações para o contêiner apropriado, garantindo uma resposta rápida e eficiente.

## Configuração das Rotas do API Gateway

O API Gateway do nosso projeto utiliza uma estrutura específica para configurar e gerenciar rotas que redirecionam solicitações para contêineres Docker específicos. Essa configuração aborda aspectos como o caminho da rota, o serviço de destino, tentativas de retentativa, checagem de saúde e mais.

Para entender completamente como configurar e o comportamento esperado dessas rotas, consulte o guia detalhado disponível em [Configuração de Rotas](./route_configuration.md).

## Configuração CORS no API Gateway

O nosso API Gateway também suporta a configuração detalhada do CORS (Cross-Origin Resource Sharing), permitindo controlar com precisão quais origens externas podem interagir com os nossos serviços e como elas podem fazer isso. A configuração aborda elementos como origens permitidas, métodos HTTP aceitos, cabeçalhos permitidos, e mais.

Para uma compreensão aprofundada de como implementar e o comportamento esperado dessas configurações CORS, consulte o guia detalhado disponível em [Configuração CORS](./cors_configuration.md).

## Inicialização do Ambiente de Desenvolvimento

Para configurar e iniciar o ambiente de desenvolvimento do projeto, consulte o guia [Desenvolvimento](./development.md).

## Inicialização do Ambiente de Produção

Para configurar e iniciar o ambiente de produção do projeto, consulte o guia [Produção](./production.md).

## Como Contribuir

Estamos sempre abertos a contribuições! Se você deseja ajudar a melhorar o projeto, seja através de correções de bugs, melhorias ou novas funcionalidades, siga nosso [Guia de Contribuição](contributing.md) para entender o processo e garantir que sua contribuição seja integrada da melhor forma possível.

## Código de Conduta

Estamos comprometidos em proporcionar uma comunidade acolhedora e inclusiva para todos. Esperamos que todos os participantes do projeto sigam nosso [Código de Conduta](code_of_coduct.md). Pedimos que leia e siga estas diretrizes para garantir um ambiente respeitoso e produtivo para todos os colaboradores.

## Licença

Este projeto está licenciado sob a licença Apache 2.0. Consulte o arquivo [LICENSE](license) para obter detalhes.
