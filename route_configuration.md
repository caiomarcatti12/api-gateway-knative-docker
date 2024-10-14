# Configuração de Hosts e Rotas do API Gateway

Este documento descreve a nova estrutura de configuração para o nosso API Gateway, que gerencia o redirecionamento de solicitações para contêineres Docker específicos e monitora sua saúde e ciclo de vida. A configuração foi atualizada para incluir **hosts** e **CORS**.

---

## Configuração YAML

A estrutura básica do arquivo de configuração é a seguinte:

```yaml
- host: host.docker.internal
  cors:
    allowedOrigins:
      - "http://host.docker.internal"
    allowedMethods:
      - "GET"
      - "POST"
      - "PUT"
      - "DELETE"
    allowedHeaders:
      - "Authorization"
      - "Content-Type"
    allowCredentials: true
    exposedHeaders:
      - "X-Custom-Header"
    maxAge: 3600
  routes:
    - path: /my-app-route
      stripPath: true
      ttl: 3
      backend:
        protocol: "http"
        host: "host.docker.internal"
        port: 8002
        containerName: "my-app-container-name"
      retry:
        attempts: 3
        period: 5
      livenessProbe:
        path: healthcheck
        successThreshold: 1
        initialDelaySeconds: 3
```

---

## Descrição dos Parâmetros

### **HostConfig** (Configuração do Host)
1. **host**: Define o domínio ou IP do host onde as rotas serão configuradas.
2. **cors**: Configura as permissões de **CORS** (Cross-Origin Resource Sharing):
   - **allowedOrigins**: Lista de origens permitidas.
   - **allowedMethods**: Métodos HTTP permitidos (GET, POST, PUT, DELETE, etc.).
   - **allowedHeaders**: Cabeçalhos permitidos nas solicitações.
   - **allowCredentials**: Define se as credenciais são permitidas.
   - **exposedHeaders**: Lista de cabeçalhos que podem ser expostos ao cliente.
   - **maxAge**: Tempo máximo, em segundos, que uma resposta CORS pode ser armazenada em cache.

### **RouteConfig** (Configuração da Rota)
1. **path**: Define o caminho da rota para redirecionamento de solicitações.
2. **stripPath**: Indica se o caminho da solicitação deve ser removido antes do redirecionamento.
3. **ttl**: Define o tempo máximo de inatividade, em segundos, para encerrar o contêiner.
4. **backend**: Contém a configuração do serviço de backend:
   - **protocol**: Protocolo usado (http ou https).
   - **host**: Host ou domínio do serviço de backend.
   - **port**: Porta onde o serviço está escutando.
   - **containerName**: Nome do contêiner correspondente.
5. **retry**: Configura as tentativas de nova tentativa para serviços indisponíveis:
   - **attempts**: Número máximo de tentativas.
   - **period**: Intervalo, em segundos, entre as tentativas.
6. **livenessProbe**: Configura a checagem de saúde do serviço:
   - **path**: Caminho para a checagem de saúde (healthcheck).
   - **successThreshold**: Número mínimo de checagens bem-sucedidas para considerar o serviço saudável.
   - **initialDelaySeconds**: Tempo inicial de espera antes da primeira checagem.

---

## Comportamento Baseado na Configuração

- **Hosts e Rotas**:  
  O API Gateway redireciona solicitações com base na configuração de **hosts**. Cada host pode ter múltiplas rotas configuradas.

- **CORS**:  
  As permissões CORS são configuradas para cada host, garantindo que apenas origens e métodos permitidos possam acessar os recursos.

- **Redirecionamento de Rotas**:  
  Quando uma solicitação chega em um caminho específico, ela é encaminhada para o backend definido na configuração da rota.

- **Checagem de Saúde**:  
  Durante a inicialização, o API Gateway realiza checagens de saúde no caminho especificado (`livenessProbe.path`).  
  - Se a checagem for bem-sucedida dentro do número de tentativas permitido (`successThreshold`), o contêiner é considerado saudável.
  - Caso contrário, o sistema tentará novamente com base nas configurações de `retry`.

- **TTL (Time To Live)**:  
  Se o contêiner não receber novas solicitações dentro do tempo configurado (`ttl`), ele será finalizado.

- **Retry**:  
  Se o contêiner falhar durante a inicialização ou se tornar inacessível, o API Gateway realizará as tentativas de nova tentativa de acordo com o número e período definidos em `retry`.

---

## Exemplo

```yaml
- host: host.docker.internal
  cors:
    allowedOrigins:
      - "http://host.docker.internal"
    allowedMethods:
      - "GET"
      - "POST"
      - "PUT"
      - "DELETE"
    allowedHeaders:
      - "Authorization"
      - "Content-Type"
    allowCredentials: true
    exposedHeaders:
      - "X-Custom-Header"
    maxAge: 3600
  routes:
    - path: /my-app-route
      stripPath: true
      ttl: 3
      backend:
        protocol: "http"
        host: "host.docker.internal"
        port: 8002
        containerName: "my-app-container-name"
      retry:
        attempts: 3
        period: 5
      livenessProbe:
        path: healthcheck
        successThreshold: 1
        initialDelaySeconds: 3
```

Neste exemplo:
1. O host `host.docker.internal` possui uma rota em `/my-app-route` que redireciona para o serviço `my-app-container-name`.
2. O backend usa o protocolo `http` na porta `8002`.
3. O contêiner será encerrado após `3` segundos de inatividade (TTL).
4. Se o serviço falhar, o API Gateway tentará reconectá-lo `3` vezes, com um intervalo de `5` segundos entre as tentativas.
5. A checagem de saúde será feita no caminho `/healthcheck` com um **delay inicial** de 3 segundos e uma **tolerância de 1 sucesso** para considerá-lo saudável.
6. O CORS está configurado para permitir origens e métodos específicos, com cache de respostas por até 3600 segundos.
