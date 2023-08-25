# Configuração de Rotas do API Gateway

Este documento descreve a configuração das rotas para o nosso API Gateway, que controla o redirecionamento de solicitações para contêineres Docker específicos e gerencia sua inicialização e saúde.

## Configuração YAML

A estrutura básica do arquivo de configuração é a seguinte:

```yaml
routes:
  - path: CAMINHO_DA_ROTA
    protocol: PROTOCOLO (http ou https)
    host: DOMÍNIO_OU_IP
    stripPath: REMOVE_PATH_DA_ROTA
    service: NOME_DO_SERVIÇO
    port: PORTA
    ttl: TEMPO_DE_VIDA
    retry: TENTATIVAS_DE_RETRY
    retryDelay: INTERVALO_ENTRE_RETRY
    healthPath: CAMINHO_DA_CHECAGEM_DE_SAÚDE
```

### Descrição dos Parâmetros

1. **path**: Define o caminho da rota na solicitação que será redirecionada pelo API Gateway.
2. **protocol**: Especifica o protocolo da solicitação (pode ser "http" ou "https").
3. **host**: Define o domínio ou IP para o qual a solicitação será redirecionada.
3. **stripPath**: Define se a solicitação que será redirecionada irá remover o parametro path da rota.
4. **service**: Especifica o nome do serviço de destino que corresponde a um contêiner Docker.
5. **port**: Indica a porta do contêiner de destino que receberá a solicitação.
6. **ttl**: Define o tempo máximo de vida do contêiner desde o momento da inicialização, medido em segundos.
7. **retry**: Estipula o número de tentativas que o API Gateway fará para iniciar o contêiner se ele não estiver funcionando após a inicialização.
8. **retryDelay**: Define o intervalo, em segundos, entre tentativas consecutivas de retentativa.
9. **healthPath**: Especifica o caminho no serviço para a checagem de saúde durante a inicialização.

## Comportamento Baseado na Configuração

- Quando uma solicitação chega ao caminho especificado (`path`), o API Gateway tenta redirecioná-la para o contêiner Docker correspondente ao serviço (`service`), usando o protocolo e host especificados, e porta (`port`) especificados.

- Se o contêiner não estiver funcionando, o API Gateway iniciará o processo de inicialização do container correspondente ao serviço (`service`) e verificará a saúde do contêiner no caminho especificado (`healthPath`).
  - Se a checagem de saúde for bem-sucedida, o API Gateway começará a redirecionar as solicitações.
  - Caso contrário, será realizado a checagem de saúde novamente até ter sucesso baseado no número especificado (`retry`). Cada tentativa ocorrerá após um intervalo de tempo (`retryDelay`).

- O contêiner tem um tempo máximo de vida (`ttl`). Se o contêiner não receber uma nova solicitação dentro desse período, ele será automaticamente encerrado.

- Se, após todas as tentativas de retentativa, o contêiner ainda não estiver saudável, o API Gateway não redirecionará as solicitações para ele, considerando-o inapto.

## Exemplo

Considere a seguinte configuração:

```yaml
routes:
  - path: "/header"
    protocol: "http"
    host: "example.com"
    service: "meu-nginx"
    port: 8081
    ttl: 600
    retry: 3
    retryDelay: 5
    healthPath: /health
```

Neste exemplo:

- Solicitações para o caminho `/header` serão redirecionadas para o contêiner Docker chamado `meu-nginx` no domínio `example.com` usando o protocolo `http`.

- O contêiner estará escutando na porta `8081`.

- Ele terá um tempo máximo de vida de `600` segundos ou `10` minutos.

- O API Gateway fará até `3` tentativas de retry com um intervalo de `5` segundos entre elas, se o contêiner não estiver inicialmente funcionando.

- Durante as tentativas de retentativa, a saúde do contêiner será verificada no caminho `/health`.

