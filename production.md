## Inicialização do Ambiente de Produção

Para configurar e iniciar o ambiente de produção do projeto "API Gateway Knative-Docker", siga os passos abaixo:

### 1. Pré-requisitos:
- Certifique-se de ter o [Docker](https://www.docker.com/get-started) e o [Docker Compose](https://docs.docker.com/compose/install/) instalados em sua máquina.

### 2. Clonar o Repositório:
Clone o repositório para sua máquina local usando o comando:
```bash
git clone https://github.com/caiomarcatti12/api-gateway-knative-docker.git
cd api-gateway-knative-docker
```

### 3. Configuração das Rotas:
- Antes de iniciar o projeto, é necessário configurar as rotas. Para isso, crie um arquivo chamado `config.yaml` na raiz do projeto.
- Use o arquivo `config.example.yaml` como referência para a estrutura do arquivo `config.yaml`.
- Configure as rotas conforme sua necessidade. Se precisar adicionar mais rotas ou entender a configuração das existentes, consulte o guia [Configuração de Rotas](./route_configuration.md).

### 4. Construção da Imagem de Produção:
- O projeto utiliza um `Dockerfile` multi-stage, que primeiro compila o código-fonte e depois cria uma imagem de produção leve usando Alpine.
- Para construir a imagem de produção, execute o seguinte comando:
```bash
docker-compose -f docker-compose-prod.yaml build
```

### 5. Iniciar o Projeto em Modo de Produção:
- Com a imagem de produção construída e o arquivo `config.yaml` configurado, você pode iniciar o projeto em modo de produção usando o Docker Compose. Execute o seguinte comando:
```bash
docker-compose -f docker-compose-prod.yaml up -d
```
Este comando iniciará o serviço de produção em modo "detached", ou seja, em segundo plano.

### 6. Verificação:
- Após iniciar o serviço, você pode verificar se o contêiner está em execução usando o comando:
```bash
docker ps
```
- Agora, o projeto deve estar em execução e pronto para receber solicitações conforme as rotas configuradas.

### 7. Parar o Projeto:
- Quando quiser parar o projeto, use o seguinte comando:
```bash
docker-compose -f docker-compose-prod.yaml down
```
